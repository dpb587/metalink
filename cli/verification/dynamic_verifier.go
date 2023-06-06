package verification

import (
	"bytes"
	"fmt"

	boshsys "github.com/cloudfoundry/bosh-utils/system"
	"github.com/dpb587/metalink"
	"github.com/dpb587/metalink/verification"
	"github.com/dpb587/metalink/verification/hash"
	"github.com/dpb587/metalink/verification/signature"
	"github.com/pkg/errors"
)

type DynamicVerifierImpl struct {
	fs boshsys.FileSystem
}

var _ DynamicVerifier = DynamicVerifierImpl{}

func NewDynamicVerifierImpl(fs boshsys.FileSystem) DynamicVerifierImpl {
	return DynamicVerifierImpl{
		fs: fs,
	}
}

func (v DynamicVerifierImpl) GetVerifier(meta4file metalink.File, skipHash bool, skipSignature bool, signatureTrustStore string) (verification.Verifier, error) {
	verifiers := []verification.Verifier{}

	if !skipHash && len(meta4file.Hashes) > 0 {
		verifiers = append(verifiers, hash.StrongestSignerVerifier)
	}

	if !skipSignature && meta4file.Signature != nil {
		if meta4file.Signature.MediaType == "application/pgp-signature" {
			trustStore, err := v.fs.ReadFile(signatureTrustStore)
			if err != nil {
				return nil, errors.Wrap(err, "Reading trust store")
			}

			verifiers = append(verifiers, signature.NewPGPVerifier(bytes.NewReader(trustStore)))
		} else {
			return nil, fmt.Errorf("Unsupported signature: %s", meta4file.Signature.MediaType)
		}
	}

	return verification.MultiVerifier{Verifiers: verifiers}, nil
}
