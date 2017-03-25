package v

import (
	"github.com/Masterminds/semver"
	"github.com/dpb587/metalink/repository"
	"github.com/dpb587/metalink/repository/sorter"
)

type Sorter struct {
	Field string
}

var _ sorter.Sorter = Sorter{}

func (s Sorter) Less(a, b repository.File) bool {
	var av, bv *semver.Version
	var err error

	av, err = semver.NewVersion(a.File.Version)
	if err != nil {
		return false
	}

	bv, err = semver.NewVersion(b.File.Version)
	if err != nil {
		return true
	}

	return !av.LessThan(bv)
}
