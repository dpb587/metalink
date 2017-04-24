package ftp

import (
	"errors"
	"fmt"
	"io"
	"net/url"
	"path/filepath"
	"time"

	"github.com/cheggaaa/pb"
	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	"github.com/dpb587/metalink/file"
	"github.com/jlaffaye/ftp"
)

type Reference struct {
	url *url.URL
}

var _ file.Reference = Reference{}

func NewReference(url *url.URL) Reference {
	return Reference{
		url: url,
	}
}

func (o Reference) Name() (string, error) {
	return filepath.Base(o.url.Path), nil
}

func (o Reference) Size() (uint64, error) {
	// @todo
	return 0, errors.New("Unsupported")
}

func (o Reference) Reader() (io.ReadCloser, error) {
	srv, err := o.connect()
	if err != nil {
		return nil, bosherr.WrapError(err, "Connecting to server")
	}

	return srv.Retr(o.url.Path)
}

func (o Reference) ReaderURI() string {
	return o.url.String()
}

func (o Reference) WriteFrom(r file.Reference, _ *pb.ProgressBar) error {
	srv, err := o.connect()
	if err != nil {
		return bosherr.WrapError(err, "Connecting to server")
	}

	reader, err := r.Reader()
	if err != nil {
		return bosherr.WrapError(err, "Opening origin for reading")
	}

	return srv.Stor(o.url.Path, reader)
}

func (o Reference) connect() (*ftp.ServerConn, error) {
	port := o.url.Port()

	if port == "" {
		port = "21"
	}

	srv, err := ftp.DialTimeout(fmt.Sprintf("%s:%s", o.url.Hostname(), port), 15*time.Second)
	if err != nil {
		return nil, bosherr.WrapError(err, "Connecting to server")
	}

	if o.url.User != nil {
		password, _ := o.url.User.Password()
		err = srv.Login(o.url.User.Username(), password)
	} else {
		err = srv.Login("anonymous", "anonymous")
	}

	if err != nil {
		return nil, bosherr.WrapError(err, "Logging in to server")
	}

	return srv, nil
}
