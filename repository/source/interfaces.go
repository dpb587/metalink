package source

import (
	"io"

	"github.com/dpb587/metalink/repository"
	"github.com/dpb587/metalink/repository/filter"
)

type Source interface {
	Load() error
	URI() string
	Filter(filter.Filter) ([]repository.RepositoryMetalink, error)
	Put(string, io.Reader) error
}

type Factory interface {
	Create(string, map[string]interface{}) (Source, error)
	Schemes() []string
}
