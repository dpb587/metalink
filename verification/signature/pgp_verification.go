package signature

import (
	"bytes"
	"io"

	"github.com/dpb587/metalink"
	"github.com/dpb587/metalink/file"
	"github.com/dpb587/metalink/verification"
	"github.com/pkg/errors"
	"golang.org/x/crypto/openpgp"
)

var endBlock = "-----END PGP PUBLIC KEY BLOCK-----"

type PGPVerification struct {
	trustStore io.Reader
}

var _ verification.Verification = PGPVerification{}

func NewPGPVerification(trustStore io.Reader) PGPVerification {
	return PGPVerification{
		trustStore: trustStore,
	}
}

func (v PGPVerification) Sign(actual file.Reference) (verification.Result, error) {
	return nil, errors.New("not yet supported")
}

func (v PGPVerification) Verify(actual file.Reference, expected metalink.File) error {
	signed, err := actual.Reader()
	if err != nil {
		return errors.Wrap(err, "Opening file for reading")
	}

	signature := bytes.NewReader([]byte(expected.Signature.Signature))

	parsedKeyRing, err := openpgp.ReadKeyRing(v.trustStore)
	if err != nil {
		return errors.Wrapf(err, "Reading armored key ring")
	}

	entity, err := openpgp.CheckArmoredDetachedSignature(parsedKeyRing, signed, signature)
	if err != nil {
		return errors.Wrap(err, "Verifying signature")
	}

	for range entity.Identities {
		return nil
		// return fmt.Sprintf("%s %s %s", entity.PrimaryKey.KeyIdShortString(), entity.PrimaryKey.CreationTime.Format("2006-01-02"), id.Name), nil
	}

	return errors.New("Unknown signature identity")
}
