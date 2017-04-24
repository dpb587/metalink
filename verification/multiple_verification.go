package verification

import (
	"github.com/dpb587/metalink"
	"github.com/dpb587/metalink/file"
)

type MultipleVerification struct {
	Verifications []Verification
}

func (v MultipleVerification) Sign(expected file.Reference) (Result, error) {
	results := MultipleResult{}

	for _, verification := range v.Verifications {
		result, err := verification.Sign(expected)
		if err != nil {
			return MultipleResult{}, err
		}

		results = append(results, result)
	}

	return results, nil
}

func (v MultipleVerification) Verify(expected file.Reference, meta4file metalink.File) error {
	for _, verification := range v.Verifications {
		if err := verification.Verify(expected, meta4file); err != nil {
			return err
		}
	}

	return nil
}
