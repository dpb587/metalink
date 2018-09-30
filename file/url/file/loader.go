package file

import (
	neturl "net/url"
	"path/filepath"
	"strings"

	"github.com/dpb587/metalink"
	"github.com/dpb587/metalink/file"
	"github.com/dpb587/metalink/file/url"
	"github.com/pkg/errors"
)

type Loader struct {}

var _ url.Loader = &Loader{}

func NewLoader() Loader {
	return Loader{}
}

func (f Loader) Schemes() []string {
	return []string{
		"file",
	}
}

func (f Loader) Load(source metalink.URL) (file.Reference, error) {
	parsedURI, err := neturl.Parse(source.URL)
	if err != nil {
		return nil, errors.Wrap(err, "Parsing source URI")
	}

	path := parsedURI.Path

	if !strings.HasPrefix(parsedURI.Path, "/") {
		path = filepath.Join(".", path)
	}

	path, err = filepath.Abs(path)
	if err != nil {
		return nil, errors.Wrap(err, "Expanding path")
	}

	return NewReference(path), nil
}
