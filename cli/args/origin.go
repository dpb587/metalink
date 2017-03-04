package args

import (
	"encoding/json"

	blobreceipt "github.com/dpb587/blob-receipt"

	bosherr "github.com/cloudfoundry/bosh-utils/errors"
)

type Origin struct {
	blobreceipt.BlobReceiptOrigin
}

func MustNewOrigin(value string) Origin {
	var a Origin

	err := a.UnmarshalFlag(value)
	if err != nil {
		panic(err)
	}

	return a
}

func (a *Origin) UnmarshalFlag(value string) error {
	a.BlobReceiptOrigin = blobreceipt.BlobReceiptOrigin{}

	if value[0] != '{' {
		a.BlobReceiptOrigin["uri"] = value

		return nil
	}

	err := json.Unmarshal([]byte(value), &a.BlobReceiptOrigin)
	if err != nil {
		return bosherr.WrapError(err, "Unmarshaling JSON origin")
	}

	return nil
}
