package origin

import (
	"errors"
	"fmt"
	"io"
	"path/filepath"
	"regexp"
	"time"

	"github.com/cheggaaa/pb"
	boshcry "github.com/cloudfoundry/bosh-utils/crypto"
	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	minio "github.com/minio/minio-go"
)

// http://docs.aws.amazon.com/general/latest/gr/rande.html#s3_region
var endpointRegex = regexp.MustCompile(`^s3(\.(dualstack\.)?|\-)[^\.]+\.amazonaws.com$`)

type S3 struct {
	client   *minio.Client
	endpoint string
	bucket   string
	object   string
	secure   bool
}

var _ Origin = S3{}

func CreateS3(client *minio.Client, secure bool, endpoint string, bucket string, object string) (Origin, error) {
	return S3{
		client:   client,
		secure:   secure,
		endpoint: endpoint,
		bucket:   bucket,
		object:   object,
	}, nil
}

func (o S3) Digest(algorithm boshcry.Algorithm) (boshcry.Digest, error) {
	return nil, errors.New("Unsupported")
}

func (o S3) Name() (string, error) {
	return filepath.Base(o.object), nil
}

func (o S3) Size() (uint64, error) {
	// @todo
	return 0, errors.New("Unsupported")
}

func (o S3) Time() (time.Time, error) {
	// @todo
	return time.Time{}, errors.New("Unsupported")
}

func (o S3) Reader() (io.ReadCloser, error) {
	reader, err := o.client.GetObject(o.bucket, o.object)
	if err != nil {
		return nil, bosherr.WrapError(err, "Opening for reading")
	}

	return reader, nil
}

func (o S3) ReaderURI() string {
	var scheme = "https"

	if !o.secure {
		scheme = "http"
	}

	return fmt.Sprintf("%s://%s/%s/%s", scheme, o.endpoint, o.bucket, o.object)
}

func (o S3) WriteFrom(from Origin, progress *pb.ProgressBar) error {
	reader, err := from.Reader()
	if err != nil {
		return bosherr.WrapError(err, "Opening from")
	}

	defer reader.Close()

	_, err = o.client.PutObjectWithProgress(o.bucket, o.object, reader, "application/octet-stream", progress)
	if err != nil {
		return bosherr.WrapError(err, "Uploading")
	}

	return nil
}
