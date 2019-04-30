package filterfactory

import (
	"fmt"

	"github.com/dpb587/metalink/repository/filter"
	"github.com/dpb587/metalink/repository/filter/axiom"
	"github.com/dpb587/metalink/repository/filter/fileversion"
	"github.com/dpb587/metalink/repository/filter/repositorypath"
)

type Manager struct {
	filters map[string]filter.FilterFactory
}

var _ filter.Manager = Manager{}

func NewManager() Manager {
	manager := Manager{}

	manager.filters = map[string]filter.FilterFactory{}
	manager.filters["axiom"] = axiom.Factory{}
	manager.filters["fileversion"] = fileversion.Factory{}
	manager.filters["repositorypath"] = repositorypath.Factory{}

	return manager
}

func (fm Manager) CreateFilter(filterType, filterValue string) (filter.Filter, error) {
	filterFactory, ok := fm.filters[filterType]
	if !ok {
		return nil, fmt.Errorf("Unknown filter type: %s", filterType)
	}

	return filterFactory.Create(filterValue)
}
