package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/cheggaaa/pb"
	"github.com/pkg/errors"
	"github.com/dpb587/metalink"
	cliverification "github.com/dpb587/metalink/cli/verification"
	"github.com/dpb587/metalink/file/metaurl"
	"github.com/dpb587/metalink/file/url"
	"github.com/dpb587/metalink/transfer"
	"github.com/dpb587/metalink/verification"
)

type FileDownload struct {
	Meta4File
	URLLoader     url.Loader                   `no-flag:"true"`
	MetaURLLoader metaurl.Loader               `no-flag:"true"`
	Verification  cliverification.DynamicVerifier `no-flag:"true"`

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
		return errors.Wrap(err, "Parsing download destination")
	}

	progress := pb.New64(int64(file.Size)).Set(pb.Bytes, true).SetRefreshRate(time.Second).SetWidth(80)

	verifier, err := c.Verification.GetVerifier(file, c.SkipHashVerification, c.SkipSignatureVerification, c.SignatureTrustStore)
	if err != nil {
		return errors.Wrap(err, "Preparing verification")
	}

	verificationResultReporter := verification.NewPrefixedVerificationResultReporter(os.Stdout, fmt.Sprintf("%s: ", file.Name))

	return transfer.NewVerifiedTransfer(c.MetaURLLoader, c.URLLoader, verifier).TransferFile(file, local, progress, verificationResultReporter)
}
