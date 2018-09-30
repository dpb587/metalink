package hash

import (
	"fmt"

	"github.com/dpb587/metalink"
	"github.com/dpb587/metalink/verification"
)

type genericVerification struct {
	hashType metalink.HashType
	hashData string
}

var _ verification.Verification = genericVerification{}

func NewGenericVerification(hashType metalink.HashType, hashData string) verification.Verification {
	return genericVerification{
		hashType: hashType,
		hashData: hashData,
	}
}

func (v genericVerification) Apply(meta4 *metalink.File) error {
	if hash, found := Find(*meta4, v.hashType); found {
		if hash.Hash == v.hashData {
			return nil
		}

		return fmt.Errorf("hash already exists with a different value: %s", v.hashType)
	}

	meta4.Hashes = append(meta4.Hashes, metalink.Hash{
		Type: v.hashType,
		Hash: v.hashData,
	})

	return nil
}

func (v genericVerification) Verify(meta4 metalink.File) verification.VerificationResult {
	expected, found := Find(meta4, v.hashType)
	if !found {
		return verification.NewSimpleVerificationResult(string(v.hashType), fmt.Errorf("hash not found: %s", v.hashType), "")
	}

	if v.hashData != expected.Hash {
		return verification.NewSimpleVerificationResult(
			string(v.hashType),
			fmt.Errorf("expected hash: %s", expected.Hash),
			fmt.Sprintf("incorrect hash: %s", v.hashData),
		)
	}

	return verification.NewSimpleVerificationResult(string(v.hashType), nil, "OK")
}

func (v genericVerification) Type() string {
	return string(v.hashType)
}

func (v genericVerification) Summary() string {
	return v.hashData
}
