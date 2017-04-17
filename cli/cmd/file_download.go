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
	FileVerifyCmd FileVerify           `no-flag:"true"`

	SkipVerification          bool   `long:"skip-verification" description:"Skip verification"`
	SkipHashVerification      bool   `long:"skip-hash-verification" description:"Skip hash verification after download"`
	SkipSignatureVerification bool   `long:"skip-signature-verification" description:"Skip signature verification after download"`
	SignatureTrustStore       string `long:"signature-trust-store" description:"Path to file with signature trust store"`

	Quiet bool `long:"quiet" short:"q" description:"Suppress passing digests"`

	Args FileDownloadArgs `positional-args:"true" required:"true"`
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

		c.FileVerifyCmd.Meta4File.Meta4.Metalink = c.Meta4File.Meta4.Metalink
		c.FileVerifyCmd.Meta4File.File = c.Meta4File.File
		c.FileVerifyCmd.Args.Local = c.Args.Local
		c.FileVerifyCmd.SkipHashVerification = c.SkipHashVerification
		c.FileVerifyCmd.SkipSignatureVerification = c.SkipSignatureVerification
		c.FileVerifyCmd.SignatureTrustStore = c.SignatureTrustStore
		c.FileVerifyCmd.Quiet = c.Quiet

		return c.FileVerifyCmd.Execute([]string{})
	}

	return errors.New("No origin blob available")
}
