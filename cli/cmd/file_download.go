package cmd

import (
	"errors"
	"time"

	"github.com/cheggaaa/pb"
	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	"github.com/dpb587/metalink/origin"
)

type FileDownload struct {
	Meta4File
	OriginFactory origin.OriginFactory `no-flag:"true"`
	VerifyCmd     FileVerify           `no-flag:"true"`
	Quiet         bool                 `long:"quiet" short:"q" description:"Suppress passing digests"`
	Hashes        []string             `long:"hash" description:"Specific hash type(s) to verify; or 'all'" default-mask:"strongest available"`
	NoVerify      bool                 `long:"no-verify" description:"Skip verification after download"`
	Args          FileDownloadArgs     `positional-args:"true" required:"true"`
}

type FileDownloadArgs struct {
	Local string `positional-arg-name:"PATH" description:"Path to the blob file"`
}

func (c *FileDownload) Execute(_ []string) error {
	file, err := c.Meta4File.Get()
	if err != nil {
		return err
	}

	local, err := c.OriginFactory.Create(c.Args.Local)
	if err != nil {
		return bosherr.WrapError(err, "Parsing origin destination")
	}

	progress := pb.New64(int64(file.Size)).SetUnits(pb.U_BYTES).SetRefreshRate(time.Second).SetWidth(80)
	progress.ShowPercent = false

	for _, url := range file.URLs {
		remote, err := c.OriginFactory.Create(url.URL)
		if err != nil {
			return bosherr.WrapError(err, "Parsing source blob")
		}

		progress.Start()

		err = local.WriteFrom(remote, progress)
		if err != nil {
			// continue
			return bosherr.WrapError(err, "Copying blob")
		}

		progress.Finish()

		if c.NoVerify {
			return nil
		}

		c.VerifyCmd.Meta4File.Meta4.Metalink = c.Meta4File.Meta4.Metalink
		c.VerifyCmd.Meta4File.File = c.Meta4File.File
		c.VerifyCmd.Hashes = c.Hashes
		c.VerifyCmd.Args.Local = c.Args.Local

		return c.VerifyCmd.Execute([]string{})
	}

	return errors.New("No origin blob available")
}
