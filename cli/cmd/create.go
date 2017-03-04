package cmd

import (
	blobreceipt "github.com/dpb587/blob-receipt"
	"github.com/dpb587/blob-receipt/cli/args"
	"github.com/dpb587/blob-receipt/origin"
	"github.com/dpb587/blob-receipt/storage"

	bosherr "github.com/cloudfoundry/bosh-utils/errors"
)

type Create struct {
	Name      *string         `long:"name" description:"Specific blob name"`
	Metadata  []args.KeyValue `long:"metadata" description:"Additional metadata to include"`
	Origin    []args.Origin   `long:"origin" description:"Additional origins include (URI string or JSON)"`
	Time      *args.Time      `long:"time" description:"Specific date/time of the blob"`
	Digest    []string        `long:"digest" description:"Specific digests to calculate" default:"md5" default:"sha1" default:"sha256" default:"sha512"`
	Overwrite bool            `long:"overwrite" description:"Overwrite receipt file if it already exists"`
	Args      CreateArgs      `positional-args:"true"`

	OriginFactory  origin.OriginFactory
	StorageFactory storage.StorageFactory
}

type CreateArgs struct {
	ReceiptPath string `positional-arg-name:"RECEIPT" description:"Path to the receipt file" required:"true"`
	BlobPath    string `positional-arg-name:"BLOB" description:"Path to the blob file"`
}

func (c *Create) Execute(_ []string) error {
	var receipt blobreceipt.BlobReceipt

	receiptStorage, err := c.StorageFactory.New(c.Args.ReceiptPath)
	if err != nil {
		return bosherr.WrapError(err, "Loading storage")
	}

	receiptExists, err := receiptStorage.Exists()
	if err != nil {
		return bosherr.WrapError(err, "Checking if receipt exists")
	}

	if receiptExists {
		receipt, err = receiptStorage.Get()
		if err != nil {
			return bosherr.WrapError(err, "Loading receipt")
		}
	}

	if c.Args.BlobPath != "" {
		origin, err := c.OriginFactory.New(c.Args.BlobPath)
		if err != nil {
			return bosherr.WrapError(err, "Loading origin")
		}

		err = receipt.UpdateFromOrigin(origin, c.Digest)
		if err != nil {
			return bosherr.WrapError(err, "Updating receipt from blob")
		}
	}

	if c.Name != nil {
		receipt.Name = *c.Name
	}

	if c.Time != nil {
		receipt.Time = (*c.Time).Time
	}

	for _, metadata := range c.Metadata {
		receipt.SetMetadata(metadata.Key, metadata.Value)
	}

	for _, origin := range c.Origin {
		receipt.SetOrigin(origin.BlobReceiptOrigin)
	}

	err = receiptStorage.Put(receipt)
	if err != nil {
		return bosherr.WrapError(err, "Putting receipt")
	}

	return nil
}
