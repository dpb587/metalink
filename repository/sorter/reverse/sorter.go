package reverse

import (
	"github.com/dpb587/metalink/repository"
	"github.com/dpb587/metalink/repository/sorter"
)

type Sorter struct {
	Sorter sorter.Sorter
}

var _ sorter.Sorter = Sorter{}

func (s Sorter) Less(a, b repository.File) bool {
	return !s.Sorter.Less(a, b)
}
