package hash

import "crypto/sha1"

var SHA1Verification = NewGenericVerification("sha-1", sha1.New)
