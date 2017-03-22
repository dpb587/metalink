package filter

import "github.com/dpb587/blob-receipt"

type Filter interface {
	IsTrue(blobreceipt.BlobReceipt) (bool, error)
}

type FilterFactory interface {
	Create(interface{}) (Filter, error)
}

type Manager interface {
	CreateFilter([]map[string]interface{}) (Filter, error)
}
