package cmd

import (
	"time"

	"github.com/cheggaaa/pb"
	"github.com/dpb587/metalink"
	"github.com/dpb587/metalink/file/url"
	"github.com/pkg/errors"
)

type FileUpload struct {
	Meta4File
	URLLoader url.Loader `no-flag:"true"`

	Location string         `long:"location" description:"ISO3166-1 country code for the geographical location"`
	Priority *uint          `long:"priority" description:"Priority value between 1 and 999999. Lower values indicate a higher priority."`
	Args     FileUploadArgs `positional-args:"true" required:"true"`
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

	local, err := c.URLLoader.LoadURL(metalink.URL{URL: c.Args.Local})
	if err != nil {
		return errors.Wrap(err, "Parsing origin destination")
	}

	remote, err := c.URLLoader.LoadURL(metalink.URL{URL: c.Args.Remote})
	if err != nil {
		return errors.Wrap(err, "Parsing source blob")
	}

	uri := remote.ReaderURI()

	for _, url := range file.URLs {
		if url.URL != uri {
			continue
		}

		return errors.New("URI already exists")
	}

	progress := pb.New64(int64(file.Size)).Set(pb.Bytes, true).SetRefreshRate(time.Second).SetWidth(80)
	progress.Start()

	err = remote.WriteFrom(local, progress)
	if err != nil {
		return errors.Wrap(err, "Copying blob")
	}

	progress.Finish()

	file.URLs = append(
		file.URLs,
		metalink.URL{
			Location: c.Location,
			Priority: c.Priority,
			URL:      uri,
		},
	)

	return c.Meta4File.Put(file)
}
