package defaultloader

import (
	"github.com/dpb587/metalink/file/metaurl"
	torrentmetaurl "github.com/dpb587/metalink/file/metaurl/torrent"
)

func New() metaurl.Loader {
	loader := metaurl.NewLoaderFactory()
	loader.Add(torrentmetaurl.Loader{})

	return loader
}
