package and

import (
	blobreceipt "github.com/dpb587/blob-receipt"
	"github.com/dpb587/blob-receipt/repository/filter"
)

type Filter struct {
	filters []filter.Filter
}

func NewFilter() Filter {
	return Filter{
		filters: []filter.Filter{},
	}
}

func (f *Filter) Add(add filter.Filter) {
	f.filters = append(f.filters, add)
}

func (f Filter) IsTrue(receipt blobreceipt.BlobReceipt) (bool, error) {
	for _, filter := range f.filters {
		is, err := filter.IsTrue(receipt)
		if err != nil {
			return false, err
		} else if is == false {
			return false, nil
		}
	}

	return true, nil
}
