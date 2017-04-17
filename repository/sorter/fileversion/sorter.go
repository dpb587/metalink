package v

import (
	"github.com/Masterminds/semver"
	"github.com/dpb587/metalink/repository"
	"github.com/dpb587/metalink/repository/sorter"
	"github.com/dpb587/metalink/repository/utility"
)

type Sorter struct {
	Field string
}

var _ sorter.Sorter = Sorter{}

func (s Sorter) Less(a, b repository.RepositoryMetalink) bool {
	var av, bv *semver.Version
	var err error

	av, err = semver.NewVersion(utility.RewriteSemiSemVer(a.Metalink.Files[0].Version))
	if err != nil {
		return false
	}

	bv, err = semver.NewVersion(utility.RewriteSemiSemVer(b.Metalink.Files[0].Version))
	if err != nil {
		return true
	}

	return !av.LessThan(bv)
}
