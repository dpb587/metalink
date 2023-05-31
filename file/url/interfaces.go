package url

import (
	"github.com/dpb587/metalink"
	"github.com/dpb587/metalink/file"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . Loader
type Loader interface {
	SupportsURL(metalink.URL) bool
	LoadURL(metalink.URL) (file.Reference, error)
}
