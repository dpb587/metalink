package hash

import (
	"crypto/sha512"

	"github.com/dpb587/metalink"
)

var SHA512SignerVerifier = NewGenericSignerVerifier(metalink.HashTypeSHA512, sha512.New)
