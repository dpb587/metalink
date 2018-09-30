package hash

import (
	"crypto/md5"

	"github.com/dpb587/metalink"
)

var MD5SignerVerifier = NewGenericSignerVerifier(metalink.HashTypeMD5, md5.New)
