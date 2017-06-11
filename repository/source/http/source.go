package http

import (
	"encoding/xml"
	"errors"
	"io"
	"io/ioutil"
	gohttp "net/http"

	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	"github.com/dpb587/metalink/repository"
	"github.com/dpb587/metalink/repository/filter"
	"github.com/dpb587/metalink/repository/source"
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
		return bosherr.WrapError(err, "Retrieving endpoint")
	} else if res.StatusCode != 200 {
		return bosherr.WrapErrorf(err, "HTTP Status %d", res.StatusCode)
	}

	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return bosherr.WrapError(err, "Reading response")
	}

	s.repo = repository.Repository{}

	err = xml.Unmarshal(bytes, &s.repo)
	if err != nil {
		return bosherr.WrapError(err, "Unmarshaling")
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
