package http

import (
	"net/http"

	"github.com/dpb587/metalink"
	"github.com/dpb587/metalink/file"
	"github.com/dpb587/metalink/file/url"
)

type Loader struct{}

var _ url.Loader = &Loader{}

func (f Loader) Schemes() []string {
	return []string{
		"http",
		"https",
	}
}

func (f Loader) Load(source metalink.URL) (file.Reference, error) {
	return NewReference(http.DefaultClient, source.URL), nil
}
