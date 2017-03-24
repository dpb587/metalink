package git

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/dpb587/metalink/repository/source"

	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	boshsys "github.com/cloudfoundry/bosh-utils/system"
)

var schemes = map[string]string{
	"git":       "git",
	"git+file":  "file",
	"git+http":  "http",
	"git+https": "https",
	"git+ssh":   "ssh",
}

type Factory struct {
	fs        boshsys.FileSystem
	cmdRunner boshsys.CmdRunner
}

var _ source.Factory = &Factory{}

func NewFactory(fs boshsys.FileSystem, cmdRunner boshsys.CmdRunner) Factory {
	return Factory{
		fs:        fs,
		cmdRunner: cmdRunner,
	}
}

func (f Factory) Schemes() []string {
	var schemeKeys = []string{}

	for scheme, _ := range schemes {
		schemeKeys = append(schemeKeys, scheme)
	}

	return schemeKeys
}

func (f Factory) Create(uri string) (source.Source, error) {
	parsedURI, err := url.Parse(uri)
	if err != nil {
		return nil, bosherr.WrapError(err, "Parsing source URI")
	}

	auth := ""

	if parsedURI.User != nil {
		auth = fmt.Sprintf("%s@", parsedURI.User.String())
	}

	splitpath := strings.SplitN(parsedURI.Path, "//", 2)
	gitpath := splitpath[0]
	fspath := ""

	if len(splitpath) == 2 {
		fspath = splitpath[1]
	}

	return NewSource(uri, fmt.Sprintf("%s://%s%s%s", schemes[parsedURI.Scheme], auth, parsedURI.Hostname(), gitpath), parsedURI.Fragment, fspath, f.fs, f.cmdRunner), nil
}
