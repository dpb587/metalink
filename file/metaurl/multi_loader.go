package metaurl

import (
	"github.com/dpb587/metalink"
	"github.com/dpb587/metalink/file"
)

type MultiLoader struct {
	loaders []Loader
}

var _ Loader = &MultiLoader{}

func NewMultiLoader(loaders ...Loader) *MultiLoader {
	return &MultiLoader{
		loaders: loaders,
	}
}

func (l *MultiLoader) SupportsMetaURL(source metalink.MetaURL) bool {
	for _, loader := range l.loaders {
		if loader.SupportsMetaURL(source) {
			return true
		}
	}

	return false
}

func (l *MultiLoader) LoadMetaURL(source metalink.MetaURL) (file.Reference, error) {
	for _, loader := range l.loaders {
		if !loader.SupportsMetaURL(source) {
			continue
		}

		ref, err := loader.LoadMetaURL(source)
		if err == UnsupportedMetaURLError {
			continue
		}

		return ref, err
	}

	return nil, UnsupportedMetaURLError
}

func (l *MultiLoader) Add(add Loader) {
	l.loaders = append(l.loaders, add)
}
