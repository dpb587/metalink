package git

import (
	"io/ioutil"
	"strings"

	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	"github.com/dpb587/metalink/repository"
	"github.com/dpb587/metalink/repository/filter"
	"github.com/dpb587/metalink/repository/source"
	minio "github.com/minio/minio-go"
)

type Source struct {
	rawURI string
	client *minio.Client
	secure bool
	bucket string
	prefix string

	files []repository.File
}

var _ source.Source = &Source{}

func NewSource(rawURI string, client *minio.Client, secure bool, bucket string, prefix string) *Source {
	return &Source{
		rawURI: rawURI,
		client: client,
		secure: secure,
		bucket: bucket,
		prefix: prefix,
	}
}

func (s *Source) Load() error {
	doneCh := make(chan struct{})
	defer close(doneCh)

	for object := range s.client.ListObjects(s.bucket, s.prefix, s.secure, doneCh) {
		if object.Err != nil {
			return bosherr.WrapError(object.Err, "Listing objects")
		} else if !strings.HasSuffix(object.Key, ".meta4") {
			continue
		}

		get, err := s.client.GetObject(s.bucket, object.Key)
		if err != nil {
			return bosherr.WrapError(err, "Getting object")
		}

		metalinkBytes, err := ioutil.ReadAll(get)
		if err != nil {
			return bosherr.WrapError(err, "Reading object")
		}

		results, err := source.ExplodeMetalinkBytes(
			repository.Repository{
				URI:     s.URI(),
				Path:    strings.TrimPrefix(object.Key, s.prefix),
				Version: object.ETag,
			},
			metalinkBytes,
		)
		if err != nil {
			return bosherr.WrapError(err, "Loading metalink")
		}

		s.files = append(s.files, results...)
	}

	return nil
}

func (s Source) URI() string {
	return s.rawURI
}

func (s Source) FilterFiles(filter filter.Filter) ([]repository.File, error) {
	return source.FilterFilesInMemory(s.files, filter)
}
