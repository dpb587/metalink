package metaurl

import (
	"fmt"

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

func (l *LoaderFactory) MediaTypes() []string {
	mediatypes := []string{}

	for _, factory := range l.factories {
		mediatypes = append(mediatypes, factory.MediaTypes()...)
	}

	return mediatypes
}

func (l *LoaderFactory) Load(source metalink.MetaURL) (file.Reference, error) {
	loader, ok := l.factories[source.MediaType]
	if !ok {
		return nil, fmt.Errorf("Unrecognized media type: %s", source.MediaType)
	}

	return loader.Load(source)
}

func (l *LoaderFactory) Add(add Loader) {
	for _, mediatype := range add.MediaTypes() {
		l.factories[mediatype] = add
	}
}
