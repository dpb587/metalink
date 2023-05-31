package torrent

import (
	"fmt"
	"io"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/anacrolix/missinggo"
	"github.com/anacrolix/torrent"
	"github.com/anacrolix/torrent/metainfo"
	"github.com/cheggaaa/pb"
	"github.com/dpb587/metalink/file"
	"github.com/pkg/errors"
)

type Reference struct {
	clientFactory ClientFactory
	url           string
	name          string
}

var _ file.Reference = Reference{}

func NewReference(clientFactory ClientFactory, url, name string) Reference {
	return Reference{
		clientFactory: clientFactory,
		url:           url,
		name:          name,
	}
}

func (o Reference) Name() (string, error) {
	return filepath.Base(o.name), nil
}

func (o Reference) Size() (uint64, error) {
	// @todo
	return 0, errors.New("Unsupported")
}

func (o Reference) Reader() (io.ReadCloser, error) {
	client, err := o.clientFactory()
	if err != nil {
		return nil, errors.Wrap(err, "starting torrent client")
	}

	var torrent *torrent.Torrent

	if strings.HasPrefix(o.url, "magnet:") {
		torrent, _ = client.AddMagnet(o.url)
	} else {
		response, err := http.DefaultClient.Get(o.url)
		if err != nil {
			return nil, errors.Wrap(err, "Loading torrent URL")
		}

		if response.StatusCode != 200 {
			return nil, fmt.Errorf("Unexpected response code: %d", response.StatusCode)
		}

		mi, err := metainfo.Load(response.Body)
		if err != nil {
			return nil, errors.Wrap(err, "Loading torrent")
		}

		torrent, _ = client.AddTorrent(mi)
	}

	<-torrent.GotInfo()

	file := torrent.Files()[0]

	return Reader{
		Reader: missinggo.NewSectionReadSeeker(torrent.NewReader(), file.Offset(), file.Length()),
		client: client,
	}, nil
}

func (o Reference) ReaderURI() string {
	// @todo generate magent?
	return o.url
}

func (o Reference) WriteFrom(r file.Reference, _ *pb.ProgressBar) error {
	return errors.New("unsupported")
}
