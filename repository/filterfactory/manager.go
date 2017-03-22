package filterfactory

import (
	"fmt"

	"github.com/dpb587/blob-receipt/repository/filter"
	"github.com/dpb587/blob-receipt/repository/filter/and"
	"github.com/dpb587/blob-receipt/repository/filter/axiom"
	"github.com/dpb587/blob-receipt/repository/filter/term"
	"github.com/dpb587/blob-receipt/repository/filter/v"
)

type Manager struct {
	filters map[string]filter.FilterFactory
}

var _ filter.Manager = Manager{}

func NewManager() Manager {
	manager := Manager{}

	manager.filters = map[string]filter.FilterFactory{}
	manager.filters["axiom"] = axiom.Factory{}
	manager.filters["term"] = term.Factory{}
	manager.filters["v"] = v.Factory{}

	return manager
}

func (fm Manager) CreateFilter(rawFilters []map[string]interface{}) (filter.Filter, error) {
	must := and.Filter{}

	for _, rawFilter := range rawFilters {
		if len(rawFilter) != 1 {
			return nil, fmt.Errorf("Filter must have a single key; has %d", len(rawFilter))
		}

		for rawFilterType, rawFilterValue := range rawFilter {
			filterFactory, ok := fm.filters[rawFilterType]
			if !ok {
				return nil, fmt.Errorf("Unknown filter type: %s", rawFilterType)
			}

			parsedFilter, err := filterFactory.Create(rawFilterValue)
			if err != nil {
				return nil, err
			}

			must.Add(parsedFilter)
		}
	}

	return must, nil
}
