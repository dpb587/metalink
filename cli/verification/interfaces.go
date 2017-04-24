package verification

import (
	"github.com/dpb587/metalink"
	"github.com/dpb587/metalink/verification"
)

type DynamicVerifier interface {
	GetVerifier(metalink.File, bool, bool, string) (verification.Verifier, error)
}
