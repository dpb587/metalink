package v

import (
	"github.com/Masterminds/semver"
	"github.com/dpb587/metalink/repository"
	"github.com/dpb587/metalink/repository/filter"
)

type Filter struct {
	Constraint semver.Constraints
}

var _ filter.Filter = Filter{}

func CreateFilter(version string) (Filter, error) {
	constraint, err := semver.NewConstraint(version)
	if err != nil {
		return Filter{}, err
	}

	return Filter{
		Constraint: *constraint,
	}, nil
}

func (f Filter) IsTrue(meta4 repository.RepositoryMetalink) (bool, error) {
	for _, file := range meta4.Metalink.Files {
		version, err := semver.NewVersion(file.Version)
		if err != nil {
			return false, err
		}

		if f.Constraint.Check(version) {
			return true, nil
		}
	}

	return false, nil
}
