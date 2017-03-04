package cmd

import (
	"errors"
	"fmt"
	"strings"

	"github.com/dpb587/blob-receipt/crypto"
	"github.com/dpb587/blob-receipt/origin"
	"github.com/dpb587/blob-receipt/storage"

	bosherr "github.com/cloudfoundry/bosh-utils/errors"
)

type Verify struct {
	Quiet  bool       `long:"quiet" short:"q" description:"Suppress passing digests"`
	Digest []string   `long:"digest" description:"Specific digest(s) to verify" default-mask:"strongest available"`
	Args   VerifyArgs `positional-args:"true" required:"true"`

	OriginFactory  origin.OriginFactory
	StorageFactory storage.StorageFactory
}

type VerifyArgs struct {
	ReceiptPath string `positional-arg-name:"RECEIPT" description:"Path to the receipt file"`
	BlobPath    string `positional-arg-name:"BLOB" description:"Path to the blob file"`
}

func (c *Verify) Execute(_ []string) error {
	receiptStorage, err := c.StorageFactory.New(c.Args.ReceiptPath)
	if err != nil {
		return bosherr.WrapError(err, "Loading storage")
	}

	receipt, err := receiptStorage.Get()
	if err != nil {
		return bosherr.WrapError(err, "Loading receipt")
	}

	origin, err := c.OriginFactory.New(c.Args.BlobPath)
	if err != nil {
		return bosherr.WrapError(err, "Loading origin")
	}

	algorithms := c.Digest

	if len(c.Digest) == 0 {
		algorithm, _ := crypto.GetStrongestAlgorithm(receipt.Digest.Algorithms())

		algorithms = []string{algorithm.Name()}
	} else if len(c.Digest) == 1 && c.Digest[0] == "all" {
		algorithms = receipt.Digest.Algorithms()
	}

	if len(algorithms) == 0 {
		return errors.New("Failed to find a digest")
	}

	failed := []string{}

	for _, algorithmName := range algorithms {
		expectedDigest, err := receipt.Digest.Get(algorithmName)
		if err != nil {
			if !c.Quiet {
				fmt.Println(fmt.Sprintf("FAIL\t%s\tERROR\t%s", algorithmName, err))
			}

			failed = append(failed, algorithmName)

			continue
		}

		reader, err := origin.Reader()
		if err != nil {
			return bosherr.WrapErrorf(err, "Opening origin for %s", algorithmName)
		}

		actualDigest, err := expectedDigest.Algorithm().CreateDigest(reader)
		if err != nil {
			fmt.Println(fmt.Sprintf("FAIL\t%s\tERROR\t%s", algorithmName, err))

			failed = append(failed, algorithmName)

			continue
		}

		if expectedDigest.String() == actualDigest.String() {
			if !c.Quiet {
				fmt.Println(fmt.Sprintf("OKAY\t%s\t%s", expectedDigest.Algorithm().Name(), crypto.GetDigestHash(expectedDigest)))
			}

			continue
		}

		fmt.Println(fmt.Sprintf("FAIL\t%s\t%s\t%s", expectedDigest.Algorithm().Name(), crypto.GetDigestHash(expectedDigest), crypto.GetDigestHash(actualDigest)))

		failed = append(failed, algorithmName)
	}

	if len(failed) > 0 {
		return fmt.Errorf("Failed to verify digest: %s", strings.Join(failed, ", "))
	}

	return nil
}
