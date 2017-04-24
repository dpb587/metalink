package hash

import (
	"fmt"
	"hash"
	"io"

	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	"github.com/dpb587/metalink"
	"github.com/dpb587/metalink/file"
	"github.com/dpb587/metalink/verification"
)

type hashNewer func() hash.Hash

type genericVerification struct {
	newer hashNewer
	type_ string
}

var _ Verification = genericVerification{}

func NewGenericVerification(type_ string, newer hashNewer) genericVerification {
	return genericVerification{
		newer: newer,
		type_: type_,
	}
}

func (v genericVerification) Sign(expected file.Reference) (verification.Result, error) {
	reader, err := expected.Reader()
	if err != nil {
		return nil, bosherr.WrapError(err, "Opening for signing")
	}

	hash := v.newer()

	_, err = io.Copy(hash, reader)
	if err != nil {
		return nil, bosherr.WrapError(err, "Reading for signing")
	}

	return NewResult(metalink.Hash{
		Type: v.type_,
		Hash: fmt.Sprintf("%x", hash.Sum(nil)),
	}), nil
}

func (v genericVerification) Verify(actual file.Reference, expected metalink.File) error {
	actualHash, err := v.Sign(actual)
	if err != nil {
		return bosherr.WrapError(err, "Calculating actual hash")
	}

	return actualHash.Verify(expected)
}

func (v genericVerification) Type() string {
	return v.type_
}
