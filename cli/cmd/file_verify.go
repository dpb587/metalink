package cmd

import (
	"os"

	"github.com/pkg/errors"
	"github.com/dpb587/metalink"
	cliverification "github.com/dpb587/metalink/cli/verification"
	"github.com/dpb587/metalink/verification"
	"github.com/dpb587/metalink/file/url"
)

type FileVerify struct {
	Meta4File
	FileLoader   url.Loader                   `no-flag:"true"`
	Verification cliverification.DynamicVerifier `no-flag:"true"`

	SkipHashVerification      bool   `long:"skip-hash-verification" description:"Skip hash verification after download"`
	SkipSignatureVerification bool   `long:"skip-signature-verification" description:"Skip signature verification after download"`
	SignatureTrustStore       string `long:"signature-trust-store" description:"Path to file with signature trust store"`

	Args FileVerifyArgs `positional-args:"true" required:"true"`
}

type FileVerifyArgs struct {
	Local string `positional-arg-name:"PATH" description:"Path to the blob file"`
}

func (c *FileVerify) Execute(_ []string) error {
	file, err := c.Meta4File.Get()
	if err != nil {
		return err
	}

	local, err := c.FileLoader.LoadURL(metalink.URL{URL: c.Args.Local})
	if err != nil {
		return errors.Wrap(err, "Parsing origin destination")
	}

	verifier, err := c.Verification.GetVerifier(file, c.SkipHashVerification, c.SkipSignatureVerification, c.SignatureTrustStore)
	if err != nil {
		return errors.Wrap(err, "Preparing verification")
	}

	result := verifier.Verify(local, file)

	verification.NewSimpleVerificationResultReporter(os.Stdout).ReportVerificationResult(file, result)

	return result.Error()
}
