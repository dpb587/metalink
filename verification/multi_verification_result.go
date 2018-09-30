package verification

type multiVerificationResult struct {
	results []VerificationResult
}

func (c multiVerificationResult) Verifier() string {
	return "multiple"
}

func (c multiVerificationResult) Error() error {
	// TODO multi-error?
	for _, result := range c.results {
		if result.Error() != nil {
			return result.Error()
		}
	}

	return nil
}

func (c multiVerificationResult) Confirmation() string {
	if c.Error() != nil {
		return ""
	}

	return "OK"
}

func (c multiVerificationResult) VerificationResults() []VerificationResult {
	return c.results
}

func NewMultiVerificationResult(results []VerificationResult) VerificationResult {
	return &multiVerificationResult{
		results: results,
	}
}
