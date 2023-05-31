package metaurl

import (
	"github.com/dpb587/metalink"
	"github.com/dpb587/metalink/file"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . Loader
type Loader interface {
	SupportsMetaURL(metalink.MetaURL) bool
	LoadMetaURL(metalink.MetaURL) (file.Reference, error)
}
