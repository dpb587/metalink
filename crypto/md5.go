package crypto

import (
	"crypto/md5"
	"fmt"
	"io"

	boshcry "github.com/cloudfoundry/bosh-utils/crypto"
	bosherr "github.com/cloudfoundry/bosh-utils/errors"
)

var DigestAlgorithmMD5 boshcry.Algorithm = algorithmMD5Impl{}

type algorithmMD5Impl struct{}

func (a algorithmMD5Impl) Name() string {
	return "md5"
}

func (a algorithmMD5Impl) CreateDigest(reader io.Reader) (boshcry.Digest, error) {
	hash := md5.New()

	_, err := io.Copy(hash, reader)
	if err != nil {
		return nil, bosherr.WrapError(err, "Reading file for digest")
	}

	return boshcry.NewDigest(a, fmt.Sprintf("%x", hash.Sum(nil))), nil
}
