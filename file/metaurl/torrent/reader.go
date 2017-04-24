package torrent

import (
	"io"

	"github.com/anacrolix/torrent"
)

type Reader struct {
	io.Reader

	client *torrent.Client
}

var _ io.ReadCloser = Reader{}

func (r Reader) Close() error {
	r.client.Close()

	return nil
}
