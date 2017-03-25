package source

import (
	"github.com/dpb587/metalink/repository"
	"github.com/dpb587/metalink/repository/filter"
)

type Source interface {
	Reload() error
	URI() string
	FilterFiles(filter.Filter) ([]repository.File, error)
}

type Factory interface {
	Create(string) (Source, error)
	Schemes() []string
}
