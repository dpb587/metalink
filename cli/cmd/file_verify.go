package cmd

import (
	"errors"
	"fmt"
	"io"
	"strings"

	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	boshsys "github.com/cloudfoundry/bosh-utils/system"
	"github.com/dpb587/metalink/origin"
	"github.com/dpb587/metalink/verify"
)

type FileVerify struct {
	Meta4File
	OriginFactory origin.OriginFactory `no-flag:"true"`
	FS            boshsys.FileSystem   `no-flag:"true"`
	Verifier      verify.Verifier      `no-flag:"true"`

	SkipHashVerification      bool   `long:"skip-hash-verification" description:"Skip hash verification after download"`
	SkipSignatureVerification bool   `long:"skip-signature-verification" description:"Skip signature verification after download"`
	SignatureTrustStore       string `long:"signature-trust-store" description:"Path to file with signature trust store"`

	Quiet bool `long:"quiet" short:"q" description:"Suppress passing digests"`

	Args FileVerifyArgs `positional-args:"true" required:"true"`
}

type FileVerifyArgs struct {
	Local string `positional-arg-name:"PATH" description:"Path to the blob file"`
}

func (c *FileVerify) Execute(_ []string) error {
	if c.SkipHashVerification && c.SkipSignatureVerification {
		return nil
	}

	file, err := c.Meta4File.Get()
	if err != nil {
		return err
	}

	local, err := c.OriginFactory.Create(c.Args.Local)
	if err != nil {
		return bosherr.WrapError(err, "Parsing origin destination")
	}

	var verifyResults = verify.MultipleVerifyResult{}

	if !c.SkipHashVerification && len(file.Hashes) > 0 {
		results, err := c.Verifier.VerifyHashes(file, local)
		if err != nil {
			return bosherr.WrapError(err, "Verifying hashes")
		}

		verifyResults = append(verifyResults, results...)
	}

	if !c.SkipSignatureVerification && file.Signature != nil {
		var trustStore io.Reader

		if c.SignatureTrustStore != "" {
			trustStoreString, err := c.FS.ReadFileString(c.SignatureTrustStore)
			if err != nil {
				return bosherr.WrapError(err, "Opening signature trust store")
			}

			trustStore = strings.NewReader(trustStoreString)
		}

		result, err := c.Verifier.VerifySignature(file, local, trustStore)
		if err != nil {
			return bosherr.WrapError(err, "Verifying signature")
		}

		verifyResults = append(verifyResults, result)
	}

	if len(verifyResults) == 0 {
		return errors.New("No hash or signature to verify")
	}

	for _, verifyResult := range verifyResults {
		var format string

		if verifyResult.HasError() {
			format = fmt.Sprintf("%s\t%%s\t%s", "FAIL", verifyResult.Error)
		} else if c.Quiet {
			continue
		} else {
			format = fmt.Sprintf("%s\t%%s", "OKAY")
		}

		if verifyResult.Type == "sign" {
			fmt.Println(fmt.Sprintf(format, fmt.Sprintf("%s\t%s", verifyResult.Type, verifyResult.Actual)))
		} else {
			fmt.Println(fmt.Sprintf(format, fmt.Sprintf("%s\t%s", verifyResult.Type, verifyResult.Expected)))
		}
	}

	if verifyResults.HasError() {
		return errors.New("Verification failed")
	}

	return nil
}
