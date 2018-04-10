package url

import (
	"fmt"
	"net/url"

	"github.com/pkg/errors"
	"github.com/dpb587/metalink"
	"github.com/dpb587/metalink/file"
)

type LoaderFactory struct {
	factories map[string]Loader
}

var _ Loader = &LoaderFactory{}

func NewLoaderFactory() *LoaderFactory {
	return &LoaderFactory{
		factories: map[string]Loader{},
	}
}

func (l *LoaderFactory) Schemes() []string {
	schemes := []string{}

	for _, factory := range l.factories {
		schemes = append(schemes, factory.Schemes()...)
	}

	return schemes
}

func (l *LoaderFactory) Load(source metalink.URL) (file.Reference, error) {
	parsedURI, err := url.Parse(source.URL)
	if err != nil {
		return nil, errors.Wrap(err, "Parsing URI")
	}

	loader, ok := l.factories[parsedURI.Scheme]
	if !ok {
		return nil, fmt.Errorf("Unrecognized scheme: %s", parsedURI.Scheme)
	}

	return loader.Load(source)
}

func (l *LoaderFactory) Add(add Loader) {
	for _, scheme := range add.Schemes() {
		l.factories[scheme] = add
	}
}
