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

func (f Filter) IsTrue(repositoryFile repository.File) (bool, error) {
	version, err := semver.NewVersion(repositoryFile.File.Version)
	if err != nil {
		return false, err
	}

	return f.Constraint.Check(version), nil
}
