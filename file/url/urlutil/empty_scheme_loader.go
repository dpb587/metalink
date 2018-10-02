package urlutil

import (
	neturl "net/url"

	"github.com/dpb587/metalink"
	"github.com/dpb587/metalink/file"
	"github.com/dpb587/metalink/file/url"
)

type emptySchemeLoader struct {
	loader url.Loader
}

var _ url.Loader = &emptySchemeLoader{}

func NewEmptySchemeLoader(loader url.Loader) url.Loader {
	return &emptySchemeLoader{
		loader: loader,
	}
}

func (f *emptySchemeLoader) SupportsURL(source metalink.URL) bool {
	parsedURI, err := neturl.Parse(source.URL)
	if err != nil {
		return false
	}

	return parsedURI.Scheme == ""
}

func (f *emptySchemeLoader) LoadURL(source metalink.URL) (file.Reference, error) {
	return f.loader.LoadURL(source)
}
