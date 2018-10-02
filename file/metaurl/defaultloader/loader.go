package defaultloader

import (
	"github.com/dpb587/metalink/file/metaurl"
	torrentmetaurl "github.com/dpb587/metalink/file/metaurl/torrent"
)

func New() metaurl.Loader {
	return metaurl.NewMultiLoader(
		torrentmetaurl.Loader{},
	)
}
