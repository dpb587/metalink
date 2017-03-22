package source

import (
	"github.com/dpb587/blob-receipt/repository"
	"github.com/dpb587/blob-receipt/repository/filter"
)

type Source interface {
	Reload() error
	URI() string
	FilterBlobReceipts(filter.Filter) ([]repository.BlobReceipt, error)
}

type Factory interface {
	Create(string) (Source, error)
	Schemes() []string
}
