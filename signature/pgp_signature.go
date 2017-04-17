package signature

import (
	"errors"
	"fmt"
	"io"

	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	"golang.org/x/crypto/openpgp"
)

var endBlock = "-----END PGP PUBLIC KEY BLOCK-----"

type PGPSignature struct{}

var _ Signature = PGPSignature{}

func (PGPSignature) Verify(trustStore, signed, signature io.Reader) (string, error) {
	parsedKeyRing, err := openpgp.ReadKeyRing(trustStore)
	if err != nil {
		return "", bosherr.WrapErrorf(err, "Reading armored key ring")
	}

	entity, err := openpgp.CheckArmoredDetachedSignature(parsedKeyRing, signed, signature)
	if err != nil {
		return "", bosherr.WrapError(err, "Verifying signature")
	}

	for _, id := range entity.Identities {
		return fmt.Sprintf("%s %s %s", entity.PrimaryKey.KeyIdShortString(), entity.PrimaryKey.CreationTime.Format("2006-01-02"), id.Name), nil
	}

	return "", errors.New("Unknown signature identity")
}

func (s PGPSignature) DefaultTrustStore() (io.Reader, error) {
	return nil, errors.New("no default trust store available")
}
