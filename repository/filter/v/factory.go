package v

import (
	"errors"
	"fmt"

	"github.com/Masterminds/semver"
	"github.com/dpb587/blob-receipt/repository/filter"
)

type Factory struct{}

func (Factory) Create(raw interface{}) (filter.Filter, error) {
	mapped, ok := raw.(map[string]string)
	if !ok {
		return nil, errors.New("Invalid type")
	} else if len(mapped) != 1 {
		return nil, fmt.Errorf("Unexpected size: %d", len(mapped))
	}

	for field, value := range mapped {
		constraint, err := semver.NewConstraint(value)
		if err != nil {
			return nil, err
		}

		return Filter{
			Field:      field,
			Constraint: *constraint,
		}, nil
	}

	panic("Unexpected empty filter")
}
