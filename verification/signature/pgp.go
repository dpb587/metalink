package signature

import (
	"bytes"
	"fmt"
	"io"

	"github.com/dpb587/metalink"
	"github.com/dpb587/metalink/file"
	"github.com/dpb587/metalink/verification"
	"github.com/pkg/errors"
	"golang.org/x/crypto/openpgp"
)

var endBlock = "-----END PGP PUBLIC KEY BLOCK-----"

type PGPVerifier struct {
	trustStore io.Reader
}

var _ verification.Verifier = PGPVerifier{}

func NewPGPVerifier(trustStore io.Reader) verification.Verifier {
	return PGPVerifier{
		trustStore: trustStore,
	}
}

func (v PGPVerifier) Verify(actual file.Reference, expected metalink.File) verification.VerificationResult {
	signed, err := actual.Reader()
	if err != nil {
		return verification.NewSimpleVerificationResult("pgp", errors.Wrap(err, "Opening file for reading"), "")
	}

	signature := bytes.NewReader([]byte(expected.Signature.Signature))

	parsedKeyRing, err := openpgp.ReadKeyRing(v.trustStore)
	if err != nil {
		return verification.NewSimpleVerificationResult("pgp", errors.Wrapf(err, "Reading armored key ring"), "")
	}

	entity, err := openpgp.CheckArmoredDetachedSignature(parsedKeyRing, signed, signature)
	if err != nil {
		return verification.NewSimpleVerificationResult("pgp", errors.Wrap(err, "Verifying signature"), "")
	}

	for _, identity := range entity.Identities {
		return verification.NewSimpleVerificationResult(
			"pgp",
			nil,
			fmt.Sprintf("%s %s %s", entity.PrimaryKey.KeyIdShortString(), entity.PrimaryKey.CreationTime.Format("2006-01-02"), identity.Name),
		)
	}

	return verification.NewSimpleVerificationResult("pgp", errors.New("Unknown signature identity"), "signature not trusted")
}
