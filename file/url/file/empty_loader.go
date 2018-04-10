package file

import (
	"fmt"
	neturl "net/url"

	"github.com/pkg/errors"
	"github.com/dpb587/metalink"
	"github.com/dpb587/metalink/file"
	"github.com/dpb587/metalink/file/url"
)

type EmptyLoader struct {
	loader Loader
}

var _ url.Loader = &EmptyLoader{}

func NewEmptyLoader(loader Loader) EmptyLoader {
	return EmptyLoader{
		loader: loader,
	}
}

func (f EmptyLoader) Schemes() []string {
	return []string{
		"",
	}
}

func (f EmptyLoader) Load(source metalink.URL) (file.Reference, error) {
	parsedURI, err := neturl.Parse(source.URL)
	if err != nil {
		return nil, errors.Wrap(err, "Parsing source URI")
	}

	if parsedURI.Scheme == "" {
		source = metalink.URL{URL: fmt.Sprintf("file://./%s", source.URL)}
	}

	return f.loader.Load(source)
}
