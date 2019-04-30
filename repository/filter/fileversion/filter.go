package fileversion

import (
	"github.com/Masterminds/semver"
	"github.com/dpb587/metalink/repository"
	"github.com/dpb587/metalink/repository/filter"
	"github.com/dpb587/metalink/repository/utility"
)

type Filter struct {
	Constraint semver.Constraints
}

var _ filter.Filter = Filter{}

func CreateFilter(version string) (Filter, error) {
	constraint, err := semver.NewConstraint(utility.RewriteSemiSemVer(version))
	if err != nil {
		return Filter{}, err
	}

	return Filter{
		Constraint: *constraint,
	}, nil
}

func (f Filter) IsTrue(meta4 repository.RepositoryMetalink) (bool, error) {
	for _, file := range meta4.Metalink.Files {
		if file.Version == "" {
			continue
		}

		version, err := semver.NewVersion(utility.RewriteSemiSemVer(file.Version))
		if err != nil {
			return false, err
		}

		if f.Constraint.Check(version) {
			return true, nil
		}
	}

	return false, nil
}
