package exists

import (
	"errors"
	"fmt"

	"github.com/dpb587/blob-receipt/repository/filter"
)

type Factory struct{}

func (Factory) Create(raw interface{}) (filter.Filter, error) {
	mapped, ok := raw.(map[string]interface{})
	if !ok {
		return nil, errors.New("Invalid type")
	} else if len(mapped) != 1 {
		return nil, fmt.Errorf("Unexpected size: %d", len(mapped))
	}

	for field, value := range mapped {
		if value != nil {
			return nil, fmt.Errorf("Value must be nil")
		}

		return Filter{
			Field: field,
		}, nil
	}

	panic("Unexpected empty filter")
}
