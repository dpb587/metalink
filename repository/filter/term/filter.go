package term

import (
	blobreceipt "github.com/dpb587/blob-receipt"
)

type Filter struct {
	Field string
	Value string
}

func (f Filter) IsTrue(receipt blobreceipt.BlobReceipt) (bool, error) {
	for _, metadata := range receipt.Metadata {
		if metadata.Key != f.Field {
			continue
		} else if metadata.Value != f.Value {
			continue
		}

		return true, nil
	}

	return false, nil
}
