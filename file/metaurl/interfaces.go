package metaurl

import (
	"github.com/dpb587/metalink"
	"github.com/dpb587/metalink/file"
)

//go:generate counterfeiter . Loader
type Loader interface {
	Load(metalink.MetaURL) (file.Reference, error)
	MediaTypes() []string
}
