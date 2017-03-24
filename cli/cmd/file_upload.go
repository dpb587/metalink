package cmd

import (
	"errors"
	"time"

	"github.com/cheggaaa/pb"
	blobreceipt "github.com/dpb587/blob-receipt"
	"github.com/dpb587/blob-receipt/origin"

	bosherr "github.com/cloudfoundry/bosh-utils/errors"
)

type FileUpload struct {
	Meta4File
	OriginFactory origin.OriginFactory `no-flag:"true"`
	Location      string               `long:"location" description:"ISO3166-1 country code for the geographical location"`
	Priority      uint                 `long:"priority" description:"Priority value between 1 and 999999. Lower values indicate a higher priority."`
	Args          FileUploadArgs       `positional-args:"true" required:"true"`
}

type FileUploadArgs struct {
	Local  string `positional-arg-name:"PATH" description:"Path to the blob file"`
	Remote string `positional-arg-name:"URI" description:"Origin URI for uploading"`
}

func (c *FileUpload) Execute(_ []string) error {
	file, err := c.Meta4File.Get()
	if err != nil {
		return err
	}

	local, err := c.OriginFactory.New(c.Args.Local)
	if err != nil {
		return bosherr.WrapError(err, "Parsing origin destination")
	}

	remote, err := c.OriginFactory.New(c.Args.Remote)
	if err != nil {
		return bosherr.WrapError(err, "Parsing source blob")
	}

	uri := remote.ReaderURI()

	for _, url := range file.URLs {
		if url.URL != uri {
			continue
		}

		return errors.New("URI already exists")
	}

	progress := pb.New64(int64(file.Size)).SetUnits(pb.U_BYTES).SetRefreshRate(time.Second)
	progress.Start()

	err = remote.WriteFrom(local, progress)
	if err != nil {
		return bosherr.WrapError(err, "Copying blob")
	}

	progress.Finish()

	file.URLs = append(
		file.URLs,
		blobreceipt.URL{
			Location: c.Location,
			Priority: c.Priority,
			URL:      uri,
		},
	)

	return c.Meta4File.Put(file)
}
