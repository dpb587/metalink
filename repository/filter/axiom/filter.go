package axiom

import (
	"github.com/dpb587/metalink/repository"
	"github.com/dpb587/metalink/repository/filter"
)

type Filter struct{}

var _ filter.Filter = Filter{}

func (f Filter) IsTrue(_ repository.RepositoryMetalink) (bool, error) {
	return true, nil
}
