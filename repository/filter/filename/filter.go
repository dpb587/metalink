package filename

import (
	"path/filepath"

	"github.com/dpb587/metalink/repository"
	"github.com/dpb587/metalink/repository/filter"
)

type Filter struct {
	Name string
}

var _ filter.Filter = Filter{}

func CreateFilter(name string) (Filter, error) {
	return Filter{
		Name: name,
	}, nil
}

func (f Filter) IsTrue(meta4 repository.RepositoryMetalink) (bool, error) {
	for _, file := range meta4.Metalink.Files {
		match, _ := filepath.Match(f.Name, file.Name)
		if match {
			return true, nil
		}
	}

	return false, nil
}
