package origin

import (
	"errors"
	"fmt"
	"io"
	"net/url"
	"path/filepath"
	"time"

	"github.com/cheggaaa/pb"
	boshcry "github.com/cloudfoundry/bosh-utils/crypto"
	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	"github.com/jlaffaye/ftp"
)

type FTP struct {
	url *url.URL
}

var _ Origin = FTP{}

func CreateFTP(url *url.URL) (Origin, error) {
	return FTP{
		url: url,
	}, nil
}

func (o FTP) Digest(algorithm boshcry.Algorithm) (boshcry.Digest, error) {
	return nil, errors.New("Unsupported")
}

func (o FTP) Name() (string, error) {
	return filepath.Base(o.url.Path), nil
}

func (o FTP) Size() (uint64, error) {
	// @todo
	return 0, errors.New("Unsupported")
}

func (o FTP) Reader() (io.ReadCloser, error) {
	srv, err := o.connect()
	if err != nil {
		return nil, bosherr.WrapError(err, "Connecting to server")
	}

	return srv.Retr(o.url.Path)
}

func (o FTP) ReaderURI() string {
	return o.url.String()
}

func (o FTP) WriteFrom(r Origin, _ *pb.ProgressBar) error {
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

func (o FTP) connect() (*ftp.ServerConn, error) {
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
