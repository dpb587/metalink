package git

import (
	"io"
	"io/ioutil"
	"path"
	"strings"

	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	"github.com/dpb587/metalink"
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

	metalinks []repository.RepositoryMetalink
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

	uri := s.URI()
	s.metalinks = []repository.RepositoryMetalink{}

	for object := range s.client.ListObjects(s.bucket, s.prefix, s.secure, doneCh) {
		if object.Err != nil {
			return bosherr.WrapError(object.Err, "Listing objects")
		} else if !strings.HasSuffix(object.Key, ".meta4") {
			continue
		}

		get, err := s.client.GetObject(s.bucket, object.Key, minio.GetObjectOptions{})
		if err != nil {
			return bosherr.WrapError(err, "Getting object")
		}

		metalinkBytes, err := ioutil.ReadAll(get)
		if err != nil {
			return bosherr.WrapError(err, "Reading object")
		}

		repometa4 := repository.RepositoryMetalink{
			Reference: repository.RepositoryMetalinkReference{
				Repository: uri,
				Path:       strings.TrimPrefix(object.Key, s.prefix),
				// Version: etag?
			},
		}

		err = metalink.Unmarshal(metalinkBytes, &repometa4.Metalink)
		if err != nil {
			return bosherr.WrapError(err, "Unmarshaling")
		}

		s.metalinks = append(s.metalinks, repometa4)
	}

	return nil
}

func (s Source) URI() string {
	return s.rawURI
}

func (s Source) Filter(f filter.Filter) ([]repository.RepositoryMetalink, error) {
	return source.FilterInMemory(s.metalinks, f)
}

func (s Source) Put(name string, data io.Reader) error {
	_, err := s.client.PutObject(s.bucket, path.Join(s.prefix, name), data, 0, minio.PutObjectOptions{ContentType: "application/octet-stream"})

	return bosherr.WrapError(err, "Writing object")
}
