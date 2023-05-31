package verification

import (
	"github.com/dpb587/metalink"
	"github.com/dpb587/metalink/file"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . Signer
type Signer interface {
	Sign(file.Reference) (Verification, error)
}

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . Verifier
type Verifier interface {
	Verify(file.Reference, metalink.File) VerificationResult
}

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . SignerVerifier
type SignerVerifier interface {
	Signer
	Verifier
}

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . Verification
type Verification interface {
	Apply(*metalink.File) error
	Verify(metalink.File) VerificationResult

	Type() string
	Summary() string
}

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . VerificationResult
type VerificationResult interface {
	Verifier() string
	Error() error
	Confirmation() string
}

type MultipleVerificationResults interface {
	VerificationResults() []VerificationResult
}

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . VerificationResultReporter
type VerificationResultReporter interface {
	ReportVerificationResult(metalink.File, VerificationResult) error
}
