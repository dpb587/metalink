package source

import (
	"github.com/dpb587/metalink/repository"
	"github.com/dpb587/metalink/repository/filter"
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
