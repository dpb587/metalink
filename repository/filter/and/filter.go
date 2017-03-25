package and

import (
	"github.com/dpb587/metalink/repository"
	"github.com/dpb587/metalink/repository/filter"
)

type Filter struct {
	filters []filter.Filter
}

func NewFilter() Filter {
	return Filter{
		filters: []filter.Filter{},
	}
}

func (f *Filter) Add(add filter.Filter) {
	f.filters = append(f.filters, add)
}

var _ filter.Filter = Filter{}

func (f Filter) IsTrue(repositoryFile repository.File) (bool, error) {
	for _, filter := range f.filters {
		is, err := filter.IsTrue(repositoryFile)
		if err != nil {
			return false, err
		} else if is == false {
			return false, nil
		}
	}

	return true, nil
}
