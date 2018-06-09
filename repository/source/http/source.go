package http

import (
	"encoding/xml"
	"io"
	"io/ioutil"
	gohttp "net/http"

	"github.com/dpb587/metalink/repository"
	"github.com/dpb587/metalink/repository/filter"
	"github.com/dpb587/metalink/repository/source"
	"github.com/pkg/errors"
)

type Source struct {
	uri    string
	client *gohttp.Client

	repo repository.Repository
}

var _ source.Source = &Source{}

func NewSource(uri string, client *gohttp.Client) *Source {
	return &Source{
		uri:    uri,
		client: client,
	}
}

func (s *Source) Load() error {
	res, err := s.client.Get(s.uri)
	if err != nil {
		return errors.Wrap(err, "Retrieving endpoint")
	} else if res.StatusCode != 200 {
		return errors.Wrapf(err, "HTTP Status %d", res.StatusCode)
	}

	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return errors.Wrap(err, "Reading response")
	}

	s.repo = repository.Repository{}

	err = xml.Unmarshal(bytes, &s.repo)
	if err != nil {
		return errors.Wrap(err, "Unmarshaling")
	}

	return nil
}

func (s Source) URI() string {
	return s.uri
}

func (s Source) Filter(f filter.Filter) ([]repository.RepositoryMetalink, error) {
	return source.FilterInMemory(s.repo.Metalinks, f)
}

func (s Source) Put(_ string, _ io.Reader) error {
	return errors.New("Put is not supported")
}
