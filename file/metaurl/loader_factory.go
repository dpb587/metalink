package metaurl

import (
	"github.com/dpb587/metalink"
	"github.com/dpb587/metalink/file"
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

func (l *LoaderFactory) MediaTypes() []string {
	mediatypes := []string{}

	for _, factories := range l.factories {
		for _, factory := range factories {
			mediatypes = append(mediatypes, factory.MediaTypes()...)
		}
	}

	return mediatypes
}

func (l *LoaderFactory) Load(source metalink.MetaURL) (file.Reference, error) {
	factories, ok := l.factories[source.MediaType]
	if !ok {
		return nil, UnsupportedMetaURLError
	}

	for _, factory := range factories {
		ref, err := factory.Load(source)
		if err == UnsupportedMetaURLError {
			continue
		}

		return ref, err
	}

	return nil, UnsupportedMetaURLError
}

func (l *LoaderFactory) Add(add Loader) {
	for _, mediatype := range add.MediaTypes() {
		l.factories[mediatype] = append(l.factories[mediatype], add)
	}
}
