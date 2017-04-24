package hash

import "github.com/dpb587/metalink/verification"

type Verification interface {
	verification.Verification

	Type() string
}
