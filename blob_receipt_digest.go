package blobreceipt

import (
	"fmt"

	"github.com/dpb587/blob-receipt/crypto"

	boshcry "github.com/cloudfoundry/bosh-utils/crypto"
	bosherr "github.com/cloudfoundry/bosh-utils/errors"
)

type BlobReceiptDigest map[string]string

func (brd BlobReceiptDigest) Algorithms() []string {
	algorithms := []string{}

	for algorithm, _ := range brd {
		algorithms = append(algorithms, algorithm)
	}

	return algorithms
}

func (brd BlobReceiptDigest) Get(algorithmName string) (boshcry.Digest, error) {
	digest, found := brd[algorithmName]
	if !found {
		return nil, fmt.Errorf("Unknown digest: %s", algorithmName)
	}

	algorithm, err := crypto.GetAlgorithm(algorithmName)
	if err != nil {
		return nil, bosherr.WrapError(err, "Getting algorithm")
	}

	return boshcry.NewDigest(algorithm, digest), nil
}
