package cmd

import (
	"time"

	"github.com/cheggaaa/pb"
	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	"github.com/dpb587/metalink"
	"github.com/dpb587/metalink/cli/verification"
	"github.com/dpb587/metalink/file/metaurl"
	"github.com/dpb587/metalink/file/url"
	"github.com/dpb587/metalink/transfer"
)

type FileDownload struct {
	Meta4File
	URLLoader     url.Loader                   `no-flag:"true"`
	MetaURLLoader metaurl.Loader               `no-flag:"true"`
	Verification  verification.DynamicVerifier `no-flag:"true"`

	SkipHashVerification      bool   `long:"skip-hash-verification" description:"Skip hash verification after download"`
	SkipSignatureVerification bool   `long:"skip-signature-verification" description:"Skip signature verification after download"`
	SignatureTrustStore       string `long:"signature-trust-store" description:"Path to file with signature trust store"`

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

	local, err := c.URLLoader.Load(metalink.URL{URL: c.Args.Local})
	if err != nil {
		return bosherr.WrapError(err, "Parsing download destination")
	}

	progress := pb.New64(int64(file.Size)).Set(pb.Bytes, true).SetRefreshRate(time.Second).SetWidth(80)

	verifier, err := c.Verification.GetVerifier(file, c.SkipHashVerification, c.SkipSignatureVerification, c.SignatureTrustStore)
	if err != nil {
		return bosherr.WrapError(err, "Preparing verification")
	}

	return transfer.NewVerifiedTransfer(c.MetaURLLoader, c.URLLoader, verifier).TransferFile(file, local, progress)
}
