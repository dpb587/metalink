package hash

import (
	"fmt"
	"hash"
	"io"

	"github.com/dpb587/metalink"
	"github.com/dpb587/metalink/file"
	"github.com/dpb587/metalink/verification"
	"github.com/pkg/errors"
)

type hashNewer func() hash.Hash

type genericSignerVerifier struct {
	newer    hashNewer
	hashType metalink.HashType
}

var _ verification.SignerVerifier = genericSignerVerifier{}

func NewGenericSignerVerifier(hashType metalink.HashType, newer hashNewer) genericSignerVerifier {
	return genericSignerVerifier{
		newer:    newer,
		hashType: hashType,
	}
}

func (v genericSignerVerifier) Sign(expected file.Reference) (verification.Verification, error) {
	reader, err := expected.Reader()
	if err != nil {
		return nil, errors.Wrap(err, "Opening for signing")
	}

	hash := v.newer()

	_, err = io.Copy(hash, reader)
	if err != nil {
		return nil, errors.Wrap(err, "Reading for signing")
	}

	return NewGenericVerification(v.hashType, fmt.Sprintf("%x", hash.Sum(nil))), nil
}

func (v genericSignerVerifier) Verify(actual file.Reference, expected metalink.File) verification.VerificationResult {
	actualHash, err := v.Sign(actual)
	if err != nil {
		return verification.NewSimpleVerificationResult(string(v.hashType), errors.Wrap(err, "Calculating actual hash"), "")
	}

	return actualHash.Verify(expected)
}

func (v genericSignerVerifier) HashType() metalink.HashType {
	return v.hashType
}
