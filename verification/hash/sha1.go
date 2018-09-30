package hash

import (
	"crypto/sha1"

	"github.com/dpb587/metalink"
)

var SHA1SignerVerifier = NewGenericSignerVerifier(metalink.HashTypeSHA1, sha1.New)
