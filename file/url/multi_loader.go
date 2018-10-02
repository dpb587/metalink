package url

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

func (l *MultiLoader) SupportsURL(source metalink.URL) bool {
	for _, loader := range l.loaders {
		if loader.SupportsURL(source) {
			return true
		}
	}

	return false
}

func (l *MultiLoader) LoadURL(source metalink.URL) (file.Reference, error) {
	for _, loader := range l.loaders {
		if !loader.SupportsURL(source) {
			continue
		}

		ref, err := loader.LoadURL(source)
		if err == UnsupportedURLError {
			continue
		}

		return ref, err
	}

	return nil, UnsupportedURLError
}

func (l *MultiLoader) Add(add Loader) {
	l.loaders = append(l.loaders, add)
}
