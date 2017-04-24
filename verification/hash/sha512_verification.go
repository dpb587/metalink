package hash

import "crypto/sha512"

var SHA512Verification = NewGenericVerification("sha-512", sha512.New)
