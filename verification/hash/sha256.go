package hash

import (
	"crypto/sha256"

	"github.com/dpb587/metalink"
)

var SHA256SignerVerifier = NewGenericSignerVerifier(metalink.HashTypeSHA256, sha256.New)
