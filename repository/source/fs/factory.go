package fs

import (
	"net/url"
	"path/filepath"

	"github.com/dpb587/metalink/repository/source"

	"github.com/pkg/errors"
	boshsys "github.com/cloudfoundry/bosh-utils/system"
)

type Factory struct {
	fs boshsys.FileSystem
}

var _ source.Factory = &Factory{}

func NewFactory(fs boshsys.FileSystem) Factory {
	return Factory{
		fs: fs,
	}
}

func (f Factory) Schemes() []string {
	return []string{
		"file",
	}
}

func (f Factory) Create(uri string, _ map[string]interface{}) (source.Source, error) {
	parsedURI, err := url.Parse(uri)
	if err != nil {
		return nil, errors.Wrap(err, "Parsing source URI")
	}

	path := parsedURI.Path

	// hacky to support relative paths via file://./relative/to/cwd
	if parsedURI.Host == "." {
		path = filepath.Join(".", path)
	}

	path, err = f.fs.ExpandPath(path)
	if err != nil {
		return nil, errors.Wrap(err, "Expanding path")
	}

	return NewSource(uri, f.fs, path), nil
}
