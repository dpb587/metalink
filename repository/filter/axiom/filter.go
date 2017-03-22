package axiom

import blobreceipt "github.com/dpb587/blob-receipt"

type Filter struct{}

func (Filter) IsTrue(_ blobreceipt.BlobReceipt) (bool, error) {
	return true, nil
}
