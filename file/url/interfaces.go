package url

import (
	"github.com/dpb587/metalink"
	"github.com/dpb587/metalink/file"
)

//go:generate counterfeiter . Loader
type Loader interface {
	Load(metalink.URL) (file.Reference, error)
	Schemes() []string
}
