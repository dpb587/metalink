package factory

import (
	"fmt"
	"net/url"

	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	"github.com/dpb587/metalink/repository/source"
)

type Factory struct {
	factories map[string]source.Factory
}

var _ source.Factory = &Factory{}

func NewFactory() *Factory {
	return &Factory{
		factories: map[string]source.Factory{},
	}
}

func (s *Factory) Schemes() []string {
	schemes := []string{}

	for _, factory := range s.factories {
		schemes = append(schemes, factory.Schemes()...)
	}

	return schemes
}

func (s *Factory) Create(uri string) (source.Source, error) {
	parsedURI, err := url.Parse(uri)
	if err != nil {
		return nil, bosherr.WrapError(err, "Parsing source URI")
	}

	schemeFactory, ok := s.factories[parsedURI.Scheme]
	if !ok {
		return nil, fmt.Errorf("Unrecognized source scheme: %s", parsedURI.Scheme)
	}

	return schemeFactory.Create(uri)
}

func (s *Factory) Add(add source.Factory) {
	for _, scheme := range add.Schemes() {
		s.factories[scheme] = add
	}
}
