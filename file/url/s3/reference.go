package s3

import (
	"fmt"
	"io"
	"path/filepath"

	"github.com/cheggaaa/pb"
	"github.com/dpb587/metalink/file"
	minio "github.com/minio/minio-go"
	"github.com/pkg/errors"
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
	info, err := o.client.StatObject(o.bucket, o.object, minio.StatObjectOptions{})
	if err != nil {
		return 0, errors.Wrap(err, "Getting object stat")
	}

	return uint64(info.Size), nil
}

func (o Reference) Reader() (io.ReadCloser, error) {
	reader, err := o.client.GetObject(o.bucket, o.object, minio.GetObjectOptions{})
	if err != nil {
		return nil, errors.Wrap(err, "Opening for reading")
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
		return errors.Wrap(err, "Checking size")
	}

	reader, err := from.Reader()
	if err != nil {
		return errors.Wrap(err, "Opening from")
	}

	defer reader.Close()

	proxyReader := progress.NewProxyReader(reader)

	_, err = o.client.PutObject(o.bucket, o.object, proxyReader, int64(size), minio.PutObjectOptions{ContentType: "application/octet-stream"})
	if err != nil {
		return errors.Wrap(err, "Uploading")
	}

	return nil
}
