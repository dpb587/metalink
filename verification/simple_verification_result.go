package verification

type simpleVerificationResult struct {
	verifier     string
	error_       error
	confirmation string
}

func (c simpleVerificationResult) Verifier() string {
	return c.verifier
}

func (c simpleVerificationResult) Error() error {
	return c.error_
}

func (c simpleVerificationResult) Confirmation() string {
	return c.confirmation
}

func NewSimpleVerificationResult(verifier string, error_ error, confirmation string) VerificationResult {
	return &simpleVerificationResult{
		verifier:     verifier,
		error_:       error_,
		confirmation: confirmation,
	}
}
