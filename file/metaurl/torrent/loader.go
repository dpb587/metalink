package torrent

import (
	"github.com/anacrolix/torrent"
	"github.com/dpb587/metalink"
	"github.com/dpb587/metalink/file"
	"github.com/dpb587/metalink/file/metaurl"
)

type Loader struct{}

var _ metaurl.Loader = &Loader{}

func (f Loader) MediaTypes() []string {
	return []string{
		"application/x-bittorrent",
		"torrent",
	}
}

func (f Loader) Load(source metalink.MetaURL) (file.Reference, error) {
	return NewReference(
		func() (*torrent.Client, error) {
			return torrent.NewClient(&torrent.Config{})
		},
		source.URL,
		source.Name,
	), nil
}
