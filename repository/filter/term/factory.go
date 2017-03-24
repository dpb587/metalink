package term

import (
	"errors"
	"fmt"

	"github.com/dpb587/metalink/repository/filter"
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
		if !ok {
			return nil, errors.New("Invalid value type")
		}

		return Filter{
			Field: field,
			Value: value,
		}, nil
	}

	panic("Unexpected empty filter")
}
