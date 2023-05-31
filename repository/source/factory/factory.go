package factory

import (
	"fmt"

	"github.com/dpb587/metalink/repository/source"
	"github.com/dpb587/metalink/repository/utility"
	"github.com/pkg/errors"
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

func (s *Factory) Create(uri string, options map[string]interface{}) (source.Source, error) {
	parsedURI, err := utility.ParseUriOrGitCloneArg(uri)
	if err != nil {
		return nil, errors.Wrap(err, "Parsing source URI")
	}

	schemeFactory, ok := s.factories[parsedURI.Scheme]
	if !ok {
		return nil, fmt.Errorf("Unrecognized source scheme: %s", parsedURI.Scheme)
	}

	return schemeFactory.Create(uri, options)
}

func (s *Factory) Add(add source.Factory) {
	for _, scheme := range add.Schemes() {
		s.factories[scheme] = add
	}
}
