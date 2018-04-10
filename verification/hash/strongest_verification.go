package hash

import (
	"github.com/dpb587/metalink"
	"github.com/dpb587/metalink/file"
	"github.com/dpb587/metalink/verification"
	"github.com/pkg/errors"
)

type strongestVerification struct {
	strengths []Verification
}

var StrongestVerification = strongestVerification{
	strengths: []Verification{
		SHA512Verification,
		SHA256Verification,
		SHA1Verification,
		MD5Verification,
	},
}

var _ verification.Signer = strongestVerification{}
var _ verification.Verifier = strongestVerification{}

func (v strongestVerification) Sign(expected file.Reference) (verification.Result, error) {
	return v.strengths[0].Sign(expected)
}

func (v strongestVerification) Verify(actual file.Reference, expected metalink.File) error {
	for _, hashVerification := range v.strengths {
		_, found := Find(expected, hashVerification.Type())
		if !found {
			continue
		}

		actualHash, err := hashVerification.Sign(actual)
		if err != nil {
			return errors.Wrap(err, "calculating actual hash")
		}

		return actualHash.Verify(expected)
	}

	return errors.New("no strong hash found")
}
