package hash

import (
	"github.com/dpb587/metalink"
	"github.com/dpb587/metalink/file"
	"github.com/dpb587/metalink/verification"
	"github.com/pkg/errors"
)

type strongestSignerVerifierSignerVerifier interface{
	verification.SignerVerifier
	HashType() metalink.HashType
}


type strongestSignerVerifier struct {
	strengths []strongestSignerVerifierSignerVerifier
}

var StrongestSignerVerifier = strongestSignerVerifier{
	strengths: []strongestSignerVerifierSignerVerifier{
		SHA512SignerVerifier,
		SHA256SignerVerifier,
		SHA1SignerVerifier,
		MD5SignerVerifier,
	},
}

var _ verification.Signer = strongestSignerVerifier{}
var _ verification.Verifier = strongestSignerVerifier{}

func (v strongestSignerVerifier) Sign(expected file.Reference) (verification.Verification, error) {
	return v.strengths[0].Sign(expected)
}

func (v strongestSignerVerifier) Verify(actual file.Reference, expected metalink.File) verification.VerificationResult {
	for _, hashVerification := range v.strengths {
		_, found := Find(expected, hashVerification.HashType())
		if !found {
			continue
		}

		actualHash, err := hashVerification.Sign(actual)
		if err != nil {
			return verification.NewSimpleVerificationResult(string(hashVerification.HashType()), errors.Wrap(err, "calculating actual hash"), "")
		}

		return actualHash.Verify(expected)
	}

	return verification.NewSimpleVerificationResult("checksum", errors.New("no strong hash found"), "")
}
