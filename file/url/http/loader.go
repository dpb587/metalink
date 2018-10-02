package http

import (
	"net/http"
	neturl "net/url"

	"github.com/dpb587/metalink"
	"github.com/dpb587/metalink/file"
	"github.com/dpb587/metalink/file/url"
)

type Loader struct{}

var _ url.Loader = &Loader{}

func (f Loader) SupportsURL(source metalink.URL) bool {
	parsed, err := neturl.Parse(source.URL)
	if err != nil {
		return false
	}

	return parsed.Scheme == "http" || parsed.Scheme == "https"
}

func (f Loader) LoadURL(source metalink.URL) (file.Reference, error) {
	return NewReference(http.DefaultClient, source.URL), nil
}
