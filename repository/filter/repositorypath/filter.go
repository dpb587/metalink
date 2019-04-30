package repositorypath

import (
	"path"

	"github.com/dpb587/metalink/repository"
	"github.com/dpb587/metalink/repository/filter"
)

type Filter struct {
	Glob string
}

var _ filter.Filter = Filter{}

func CreateFilter(glob string) (Filter, error) {
	_, err := path.Match(glob, "")
	if err != nil {
		return Filter{}, err
	}

	return Filter{
		Glob: glob,
	}, nil
}

func (f Filter) IsTrue(meta4 repository.RepositoryMetalink) (bool, error) {
	match, err := path.Match(f.Glob, meta4.Reference.Path)
	if err != nil {
		return false, err
	}

	return match, nil
}
