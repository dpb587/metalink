package cmd

import (
	"time"

	"github.com/cheggaaa/pb"
	blobreceipt "github.com/dpb587/blob-receipt"
	"github.com/dpb587/blob-receipt/origin"
	"github.com/dpb587/blob-receipt/storage"

	bosherr "github.com/cloudfoundry/bosh-utils/errors"
)

type Upload struct {
	SkipReceiptUpdate bool       `long:"skip-receipt-update" description:"Do not add the origin to the receipt"`
	Args              UploadArgs `positional-args:"true" required:"true"`

	OriginFactory  origin.OriginFactory   `no-flag:"true"`
	StorageFactory storage.StorageFactory `no-flag:"true"`
}

type UploadArgs struct {
	ReceiptPath string `positional-arg-name:"RECEIPT" description:"Path to the receipt file"`
	OriginPath  string `positional-arg-name:"ORIGIN" description:"Origin URI for uploading"`
	BlobPath    string `positional-arg-name:"BLOB" description:"Path to the blob file"`
}

func (c *Upload) Execute(_ []string) error {
	receiptStorage, err := c.StorageFactory.New(c.Args.ReceiptPath)
	if err != nil {
		return bosherr.WrapError(err, "Loading storage")
	}

	receipt, err := receiptStorage.Get()
	if err != nil {
		return bosherr.WrapError(err, "Loading receipt")
	}

	destinationOrigin, err := c.OriginFactory.New(c.Args.OriginPath)
	if err != nil {
		return bosherr.WrapError(err, "Parsing origin destination")
	}

	sourceOrigin, err := c.OriginFactory.New(c.Args.BlobPath)
	if err != nil {
		return bosherr.WrapError(err, "Parsing source blob")
	}

	progress := pb.New64(int64(receipt.Size)).SetUnits(pb.U_BYTES).SetRefreshRate(time.Second)
	progress.Start()

	err = destinationOrigin.WriteFrom(sourceOrigin, progress)
	if err != nil {
		// continue
		return bosherr.WrapError(err, "Copying blob")
	}

	progress.Finish()

	receipt.SetOrigin(blobreceipt.BlobReceiptOrigin{
		"uri": destinationOrigin.ReaderURI(),
	})

	err = receiptStorage.Put(receipt)
	if err != nil {
		return bosherr.WrapError(err, "Updating receipt")
	}

	return nil
}
