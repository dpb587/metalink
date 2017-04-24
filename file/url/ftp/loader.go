package ftp

import (
	neturl "net/url"

	bosherr "github.com/cloudfoundry/bosh-utils/errors"
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
		return nil, bosherr.WrapError(err, "Parsing source URI")
	}

	return NewReference(parsedURI), nil
}
