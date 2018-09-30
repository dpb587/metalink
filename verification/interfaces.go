package verification

import (
	"github.com/dpb587/metalink"
	"github.com/dpb587/metalink/file"
)

//go:generate counterfeiter . Signer
type Signer interface {
	Sign(file.Reference) (Verification, error)
}

//go:generate counterfeiter . Verifier
type Verifier interface {
	Verify(file.Reference, metalink.File) VerificationResult
}

//go:generate counterfeiter . SignerVerifier
type SignerVerifier interface {
	Signer
	Verifier
}

//go:generate counterfeiter . Verification
type Verification interface {
	Apply(*metalink.File) error
	Verify(metalink.File) VerificationResult

	Type() string
	Summary() string
}

//go:generate counterfeiter . VerificationResult
type VerificationResult interface {
	Verifier()     string
	Error()        error
	Confirmation() string
}

type MultipleVerificationResults interface {
	VerificationResults() []VerificationResult
}

//go:generate counterfeiter . VerificationResultReporter
type VerificationResultReporter interface {
	ReportVerificationResult(VerificationResult) error
}
