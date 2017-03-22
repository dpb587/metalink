package cmd

import (
	"errors"
	"fmt"
	"time"

	blobreceipt "github.com/dpb587/blob-receipt"
	"github.com/dpb587/blob-receipt/storage"

	bosherr "github.com/cloudfoundry/bosh-utils/errors"
)

type Dump struct {
	Digest    string   `long:"digest" description:"Digest to dump" default-mask:"ALGO"`
	Metadata  string   `long:"metadata" description:"Metadata key to dump" default-mask:"KEY"`
	Name      bool     `long:"name" description:"Blob name"`
	Size      bool     `long:"size" description:"Blob size"`
	Timestamp bool     `long:"timestamp" description:"Blob timestamp"`
	Args      DumpArgs `positional-args:"true"`

	StorageFactory storage.StorageFactory
}

type DumpArgs struct {
	ReceiptPath string `positional-arg-name:"RECEIPT" description:"Path to the receipt file" required:"true"`
}

func (c *Dump) Execute(_ []string) error {
	var receipt blobreceipt.BlobReceipt

	receiptStorage, err := c.StorageFactory.New(c.Args.ReceiptPath)
	if err != nil {
		return bosherr.WrapError(err, "Loading storage")
	}

	receipt, err = receiptStorage.Get()
	if err != nil {
		return bosherr.WrapError(err, "Loading receipt")
	}

	if c.Digest != "" {
		digest, ok := receipt.Digest[c.Digest]
		if !ok {
			return fmt.Errorf("Digest not found: %s", c.Digest)
		}

		fmt.Println(digest)
	} else if c.Metadata != "" {
		for _, metadata := range receipt.Metadata {
			if metadata.Key != c.Metadata {
				continue
			}

			fmt.Println(metadata.Value)

			return nil
		}

		return fmt.Errorf("Metadata not found: %s", c.Metadata)
	} else if c.Name {
		fmt.Println(receipt.Name)
	} else if c.Size {
		fmt.Println(receipt.Size)
	} else if c.Timestamp {
		fmt.Println(receipt.Time.Format(time.RFC3339))
	} else {
		return errors.New("A dump option is required")
	}

	return nil
}
