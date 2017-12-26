package filename

import "github.com/dpb587/metalink/repository/filter"

type Factory struct{}

var _ filter.FilterFactory = Factory{}

func (Factory) Create(name string) (filter.Filter, error) {
	return CreateFilter(name)
}
