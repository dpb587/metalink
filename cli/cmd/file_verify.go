package cmd

import (
	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	"github.com/dpb587/metalink"
	"github.com/dpb587/metalink/cli/verification"
	"github.com/dpb587/metalink/file/url"
)

type FileVerify struct {
	Meta4File
	FileLoader   url.Loader                   `no-flag:"true"`
	Verification verification.DynamicVerifier `no-flag:"true"`

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

	local, err := c.FileLoader.Load(metalink.URL{URL: c.Args.Local})
	if err != nil {
		return bosherr.WrapError(err, "Parsing origin destination")
	}

	verifier, err := c.Verification.GetVerifier(file, c.SkipHashVerification, c.SkipSignatureVerification, c.SignatureTrustStore)
	if err != nil {
		return bosherr.WrapError(err, "Preparing verification")
	}

	return verifier.Verify(local, file)
}
