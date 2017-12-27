package s3

import (
	"errors"
	"fmt"
	"io"
	"path/filepath"

	"github.com/cheggaaa/pb"
	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	"github.com/dpb587/metalink/file"
	minio "github.com/minio/minio-go"
)

type Reference struct {
	client   *minio.Client
	endpoint string
	bucket   string
	object   string
	secure   bool
}

var _ file.Reference = Reference{}

func NewReference(client *minio.Client, secure bool, endpoint string, bucket string, object string) Reference {
	return Reference{
		client:   client,
		secure:   secure,
		endpoint: endpoint,
		bucket:   bucket,
		object:   object,
	}
}

func (o Reference) Name() (string, error) {
	return filepath.Base(o.object), nil
}

func (o Reference) Size() (uint64, error) {
	// @todo
	return 0, errors.New("Unsupported")
}

func (o Reference) Reader() (io.ReadCloser, error) {
	reader, err := o.client.GetObject(o.bucket, o.object, minio.GetObjectOptions{})
	if err != nil {
		return nil, bosherr.WrapError(err, "Opening for reading")
	}

	return reader, nil
}

func (o Reference) ReaderURI() string {
	var scheme = "https"

	if !o.secure {
		scheme = "http"
	}

	return fmt.Sprintf("%s://%s/%s/%s", scheme, o.endpoint, o.bucket, o.object)
}

func (o Reference) WriteFrom(from file.Reference, progress *pb.ProgressBar) error {
	size, err := from.Size()
	if err != nil {
		return bosherr.WrapError(err, "Checking size")
	}

	reader, err := from.Reader()
	if err != nil {
		return bosherr.WrapError(err, "Opening from")
	}

	defer reader.Close()

	proxyReader := progress.NewProxyReader(reader)

	_, err = o.client.PutObject(o.bucket, o.object, proxyReader, int64(size), minio.PutObjectOptions{ContentType: "application/octet-stream"})
	if err != nil {
		return bosherr.WrapError(err, "Uploading")
	}

	return nil
}
