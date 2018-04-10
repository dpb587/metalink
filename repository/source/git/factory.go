package git

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/dpb587/metalink/repository/source"

	"github.com/pkg/errors"
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

func (f Factory) Create(uri string, options map[string]interface{}) (source.Source, error) {
	parsedURI, err := url.Parse(uri)
	if err != nil {
		return nil, errors.Wrap(err, "Parsing source URI")
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

	var privateKey *string
	commits := sourceCommitSettings{
		authorEmail:    "metalink-repository@localhost",
		authorName:     "metalink-repository",
		committerEmail: "metalink-repository@localhost",
		committerName:  "metalink-repository",
		message:        "update metalink",
	}

	if val, found := options["private_key"]; found {
		privateKeyQ := val.(string)
		privateKey = &privateKeyQ
	}

	if val, found := options["author_email"]; found {
		commits.authorEmail = val.(string)
	}

	if val, found := options["author_name"]; found {
		commits.authorName = val.(string)
	}

	if val, found := options["committer_email"]; found {
		commits.committerEmail = val.(string)
	}

	if val, found := options["committer_name"]; found {
		commits.committerName = val.(string)
	}

	if val, found := options["message"]; found {
		commits.message = val.(string)
	}

	return NewSource(uri, strings.TrimPrefix(fmt.Sprintf("%s://%s%s%s", schemes[parsedURI.Scheme], auth, parsedURI.Host, gitpath), "ssh://"), parsedURI.Fragment, fspath, privateKey, commits, f.fs, f.cmdRunner), nil
}
