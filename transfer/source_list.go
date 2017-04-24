package transfer

import (
	"sort"

	"github.com/dpb587/metalink"
)

type source struct {
	Priority uint
	URL      *metalink.URL
	MetaURL  *metalink.MetaURL
}

type sourceList []source

func (a sourceList) Len() int           { return len(a) }
func (a sourceList) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a sourceList) Less(i, j int) bool { return a[i].Priority < a[j].Priority }

func newSourceList(metaurls []metalink.MetaURL, urls []metalink.URL) sourceList {
	sources := sourceList{}

	for _, metaurl := range metaurls {
		var priority = uint(999999)
		if metaurl.Priority != nil {
			priority = *metaurl.Priority
		}

		sources = append(
			sources,
			source{
				MetaURL:  &metaurl,
				Priority: priority,
			},
		)
	}

	for _, url := range urls {
		var priority = uint(999999)
		if url.Priority != nil {
			priority = *url.Priority
		}

		sources = append(
			sources,
			source{
				URL:      &url,
				Priority: priority,
			},
		)
	}

	sort.Sort(sources)

	return sources
}
