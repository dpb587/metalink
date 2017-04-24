package verification

import (
	"github.com/dpb587/metalink"
	"github.com/dpb587/metalink/file"
)

//go:generate counterfeiter . Signer
type Signer interface {
	Sign(file.Reference) (Result, error)
}

//go:generate counterfeiter . Verifier
type Verifier interface {
	Verify(file.Reference, metalink.File) error
}

//go:generate counterfeiter . Verification
type Verification interface {
	Signer
	Verifier
}

//go:generate counterfeiter . Result
type Result interface {
	Apply(*metalink.File) error
	Verify(metalink.File) error

	Type() string
	Summary() string
}
