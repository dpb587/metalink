package filter

import "github.com/dpb587/metalink"

type Filter interface {
	IsTrue(metalink.BlobReceipt) (bool, error)
}

type FilterFactory interface {
	Create(interface{}) (Filter, error)
}

type Manager interface {
	CreateFilter([]map[string]interface{}) (Filter, error)
}
