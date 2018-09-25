package url

import (
	"net/url"

	"github.com/dpb587/metalink"
	"github.com/dpb587/metalink/file"
	"github.com/pkg/errors"
)

type LoaderFactory struct {
	factories map[string][]Loader
}

var _ Loader = &LoaderFactory{}

func NewLoaderFactory() *LoaderFactory {
	return &LoaderFactory{
		factories: map[string][]Loader{},
	}
}

func (l *LoaderFactory) Schemes() []string {
	schemes := []string{}

	for _, factories := range l.factories {
		for _, factory := range factories {
			schemes = append(schemes, factory.Schemes()...)
		}
	}

	return schemes
}

func (l *LoaderFactory) Load(source metalink.URL) (file.Reference, error) {
	parsedURI, err := url.Parse(source.URL)
	if err != nil {
		return nil, errors.Wrap(err, "Parsing URI")
	}

	factories, ok := l.factories[parsedURI.Scheme]
	if !ok {
		return nil, UnsupportedURLError
	}

	for _, factory := range factories {
		ref, err := factory.Load(source)
		if err == UnsupportedURLError {
			continue
		}

		return ref, err
	}

	return nil, UnsupportedURLError
}

func (l *LoaderFactory) Add(add Loader) {
	for _, scheme := range add.Schemes() {
		l.factories[scheme] = append(l.factories[scheme], add)
	}
}
