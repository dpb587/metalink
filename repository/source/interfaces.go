package source

import (
	"github.com/dpb587/metalink/repository"
	"github.com/dpb587/metalink/repository/filter"
)

type Source interface {
	Load() error
	URI() string
	Filter(filter.Filter) ([]repository.RepositoryMetalink, error)
}

type Factory interface {
	Create(string) (Source, error)
	Schemes() []string
}
