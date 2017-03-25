package filter

import "github.com/dpb587/metalink/repository"

type Filter interface {
	IsTrue(repository.File) (bool, error)
}

type FilterFactory interface {
	Create(string) (Filter, error)
}

type Manager interface {
	CreateFilter(string, string) (Filter, error)
}
