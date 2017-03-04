package args

import (
	"github.com/dpb587/blob-receipt/crypto"

	boshcry "github.com/cloudfoundry/bosh-utils/crypto"
)

type DigestAlgorithm struct {
	boshcry.Algorithm
}

func MustNewDigestAlgorithm(value string) DigestAlgorithm {
	var a DigestAlgorithm

	err := a.UnmarshalFlag(value)
	if err != nil {
		panic(err)
	}

	return a
}

func (a *DigestAlgorithm) UnmarshalFlag(value string) error {
	algorithm, err := crypto.GetAlgorithm(value)
	if err != nil {
		return err
	}

	a.Algorithm = algorithm

	return nil
}
