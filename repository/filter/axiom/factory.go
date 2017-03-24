package axiom

import "github.com/dpb587/metalink/repository/filter"

type Factory struct{}

func (Factory) Create(raw interface{}) (filter.Filter, error) {
	return Filter{}, nil
}
