package exists

import (
	blobreceipt "github.com/dpb587/blob-receipt"
)

type Filter struct {
	Field string
}

func (f Filter) IsTrue(receipt blobreceipt.BlobReceipt) (bool, error) {
	for _, metadata := range receipt.Metadata {
		if metadata.Key != f.Field {
			continue
		}

		return true, nil
	}

	return false, nil
}
