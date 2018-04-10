package file

import (
	neturl "net/url"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
	boshsys "github.com/cloudfoundry/bosh-utils/system"
	"github.com/dpb587/metalink"
	"github.com/dpb587/metalink/file"
	"github.com/dpb587/metalink/file/url"
)

type Loader struct {
	fs boshsys.FileSystem
}

var _ url.Loader = &Loader{}

func NewLoader(fs boshsys.FileSystem) Loader {
	return Loader{
		fs: fs,
	}
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

	if !strings.HasPrefix(parsedURI.Path, "//") {
		path = filepath.Join(".", path)
	}

	path, err = f.fs.ExpandPath(path)
	if err != nil {
		return nil, errors.Wrap(err, "Expanding path")
	}

	return NewReference(f.fs, path), nil
}
