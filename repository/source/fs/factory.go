package fs

import (
	"net/url"

	"github.com/dpb587/metalink/repository/source"

	bosherr "github.com/cloudfoundry/bosh-utils/errors"
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

func (f Factory) Create(uri string) (source.Source, error) {
	parsedURI, err := url.Parse(uri)
	if err != nil {
		return nil, bosherr.WrapError(err, "Parsing source URI")
	}

	return NewSource(uri, f.fs, parsedURI.Path), nil
}
