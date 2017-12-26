package filter

import "github.com/dpb587/metalink/repository"

type Filterer interface {
	Filter(repository.RepositoryMetalink) (*repository.RepositoryMetalink, error)
}

type FilterFactory interface {
	Create(string) (Filterer, error)
}
