package hash

import "crypto/sha256"

var SHA256Verification = NewGenericVerification("sha-256", sha256.New)
