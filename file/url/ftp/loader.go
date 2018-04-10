package ftp

import (
	neturl "net/url"

	"github.com/pkg/errors"
	"github.com/dpb587/metalink"
	"github.com/dpb587/metalink/file"
	"github.com/dpb587/metalink/file/url"
)

type Loader struct{}

var _ url.Loader = &Loader{}

func (f Loader) Schemes() []string {
	return []string{
		"ftp",
	}
}

func (f Loader) Load(source metalink.URL) (file.Reference, error) {
	parsedURI, err := neturl.Parse(source.URL)
	if err != nil {
		return nil, errors.Wrap(err, "Parsing source URI")
	}

	return NewReference(parsedURI), nil
}
