package cmd

import (
	"errors"
	"time"

	"github.com/cheggaaa/pb"
	"github.com/dpb587/blob-receipt/origin"
	"github.com/dpb587/blob-receipt/storage"

	bosherr "github.com/cloudfoundry/bosh-utils/errors"
)

type Download struct {
	Origin []string     `long:"origin" description:"Specific origin(s) to use" default-mask:"automatic"`
	Args   DownloadArgs `positional-args:"true" required:"true"`

	OriginFactory  origin.OriginFactory   `no-flag:"true"`
	StorageFactory storage.StorageFactory `no-flag:"true"`

	VerifyCommand Verify `no-flag:"true"`
}

type DownloadArgs struct {
	ReceiptPath string `positional-arg-name:"RECEIPT" description:"Path to the receipt file"`
	BlobPath    string `positional-arg-name:"BLOB" description:"Path to the blob file"`
}

func (c *Download) Execute(_ []string) error {
	receiptStorage, err := c.StorageFactory.New(c.Args.ReceiptPath)
	if err != nil {
		return bosherr.WrapError(err, "Loading storage")
	}

	receipt, err := receiptStorage.Get()
	if err != nil {
		return bosherr.WrapError(err, "Loading receipt")
	}

	destinationOrigin, err := c.OriginFactory.New(c.Args.BlobPath)
	if err != nil {
		return bosherr.WrapError(err, "Parsing destination blob")
	}

	originURIs := c.Origin
	if len(originURIs) == 0 {
		for _, origin := range receipt.Origin {
			originURIs = append(originURIs, origin.URI())
		}
	}

	for _, originURI := range originURIs {
		sourceOrigin, err := c.OriginFactory.New(originURI)
		if err != nil {
			return bosherr.WrapError(err, "Parsing origin")
		}

		progress := pb.New64(int64(receipt.Size)).SetUnits(pb.U_BYTES).SetRefreshRate(time.Second)
		progress.Start()

		err = destinationOrigin.WriteFrom(sourceOrigin, progress)
		if err != nil {
			// continue
			return bosherr.WrapError(err, "Copying blob")
		}

		progress.Finish()

		c.VerifyCommand.Args.ReceiptPath = c.Args.ReceiptPath
		c.VerifyCommand.Args.BlobPath = c.Args.BlobPath

		return c.VerifyCommand.Execute([]string{})
	}

	return errors.New("No origin blob available")
}
