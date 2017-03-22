package reverse

import (
	"github.com/dpb587/blob-receipt/repository"
	"github.com/dpb587/blob-receipt/repository/sorter"
)

type Sorter struct {
	Sorter sorter.Sorter
}

var _ sorter.Sorter = Sorter{}

func (s Sorter) Less(a, b repository.BlobReceipt) bool {
	return !s.Sorter.Less(a, b)
}
