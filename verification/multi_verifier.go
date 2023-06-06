package verification

import (
	"github.com/dpb587/metalink"
	"github.com/dpb587/metalink/file"
)

type MultiVerifier struct {
	Verifiers []Verifier
}

func (v MultiVerifier) Verify(expected file.Reference, meta4file metalink.File) VerificationResult {
	var results []VerificationResult

	for _, verification := range v.Verifiers {
		result := verification.Verify(expected, meta4file)

		if result.Error() != nil {
			return result
		}

		results = append(results, result)
	}

	return NewMultiVerificationResult(results)
}
