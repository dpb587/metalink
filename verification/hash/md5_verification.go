package hash

import "crypto/md5"

var MD5Verification = NewGenericVerification("md5", md5.New)
