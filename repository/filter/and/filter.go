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

func (f Filter) IsTrue(meta4 repository.RepositoryMetalink) (bool, error) {
	for _, filter := range f.filters {
		is, err := filter.IsTrue(meta4)
		if err != nil {
			return false, err
		} else if is == false {
			return false, nil
		}
	}

	return true, nil
}
