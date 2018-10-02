package urlutil

import (
	"regexp"

	"github.com/dpb587/metalink"
	"github.com/dpb587/metalink/file"
	"github.com/dpb587/metalink/file/url"
)

type filteredLoader struct {
	loader  url.Loader
	include []*regexp.Regexp
	exclude []*regexp.Regexp
}

func NewFilteredLoader(loader url.Loader, include, exclude []*regexp.Regexp) url.Loader {
	return &filteredLoader{
		loader: loader,
		include: include,
		exclude: exclude,
	}
}

func (l *filteredLoader) SupportsURL(source metalink.URL) bool {
  for _, exclude := range l.exclude {
    if exclude.MatchString(source.URL) {
      return false
    }
  }

  for _, include := range l.include {
    if include.MatchString(source.URL) {
      return true
    }
  }

	return false
}

func (l *filteredLoader) LoadURL(source metalink.URL) (file.Reference, error) {
  return l.loader.LoadURL(source)
}
