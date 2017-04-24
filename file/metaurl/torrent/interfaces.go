package torrent

import "github.com/anacrolix/torrent"

type ClientFactory func() (*torrent.Client, error)
