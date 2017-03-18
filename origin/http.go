package origin

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path/filepath"
	"time"

	"github.com/cheggaaa/pb"
	boshcry "github.com/cloudfoundry/bosh-utils/crypto"
	bosherr "github.com/cloudfoundry/bosh-utils/errors"
)

type HTTP struct {
	client *http.Client
	url    string
}

var _ Origin = HTTP{}

func CreateHTTP(client *http.Client, url string) (Origin, error) {
	return HTTP{
		client: client,
		url:    url,
	}, nil
}

func (o HTTP) Digest(algorithm boshcry.Algorithm) (boshcry.Digest, error) {
	return nil, errors.New("Unsupported")
}

func (o HTTP) Name() (string, error) {
	parsed, err := url.Parse(o.url)
	if err != nil {
		return "", bosherr.WrapError(err, "Parsing URL")
	}

	return filepath.Base(parsed.Path), nil
}

func (o HTTP) Size() (uint64, error) {
	// @todo
	return 0, errors.New("Unsupported")
}

func (o HTTP) Time() (time.Time, error) {
	return time.Time{}, errors.New("Unsupported")
}

func (o HTTP) Reader() (io.ReadCloser, error) {
	response, err := o.client.Get(o.url)
	if err != nil {
		return nil, bosherr.WrapError(err, "Loading URL")
	}

	if response.StatusCode != 200 {
		return nil, fmt.Errorf("Unexpected response code: %d", response.StatusCode)
	}

	return response.Body, nil
}

func (o HTTP) ReaderURI() string {
	return o.url
}

func (o HTTP) WriteFrom(_ Origin, _ *pb.ProgressBar) error {
	// @todo
	return errors.New("Unsupported")
}
