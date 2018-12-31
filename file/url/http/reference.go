package http

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path/filepath"
	"strconv"

	"github.com/cheggaaa/pb"
	"github.com/dpb587/metalink/file"
	"github.com/pkg/errors"
)

type Reference struct {
	client *http.Client
	url    string
}

var _ file.Reference = Reference{}

func NewReference(client *http.Client, url string) Reference {
	return Reference{
		client: client,
		url:    url,
	}
}

func (o Reference) Name() (string, error) {
	parsed, err := url.Parse(o.url)
	if err != nil {
		return "", errors.Wrap(err, "Parsing URL")
	}

	return filepath.Base(parsed.Path), nil
}

func (o Reference) Size() (uint64, error) {
	response, err := o.client.Head(o.url)
	if err != nil {
		return 0, errors.Wrap(err, "Loading URL")
	}

	if response.StatusCode != 200 {
		return 0, fmt.Errorf("Unexpected response code: %d", response.StatusCode)
	}

	lengthString := response.Header.Get("content-length")
	if lengthString == "" {
		return 0, errors.New("Content-Length not returned")
	}

	length, err := strconv.ParseUint(lengthString, 10, 64)
	if err != nil {
		return 0, errors.Wrap(err, "Converting Content-Length to int")
	}

	return length, nil
}

func (o Reference) Reader() (io.ReadCloser, error) {
	response, err := o.client.Get(o.url)
	if err != nil {
		return nil, errors.Wrap(err, "Loading URL")
	}

	if response.StatusCode != 200 {
		return nil, fmt.Errorf("Unexpected response code: %d", response.StatusCode)
	}

	return response.Body, nil
}

func (o Reference) ReaderURI() string {
	return o.url
}

func (o Reference) WriteFrom(_ file.Reference, _ *pb.ProgressBar) error {
	// @todo
	return errors.New("Unsupported")
}
