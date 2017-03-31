package origin

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"

	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	boshsys "github.com/cloudfoundry/bosh-utils/system"
	minio "github.com/minio/minio-go"
)

type defaultFactory struct {
	fs boshsys.FileSystem
}

var _ OriginFactory = defaultFactory{}

func NewDefaultFactory(fs boshsys.FileSystem) OriginFactory {
	return defaultFactory{
		fs: fs,
	}
}

// @todo rename to Create
func (f defaultFactory) Create(uri string) (Origin, error) {
	parsed, err := url.Parse(uri)
	if err != nil {
		return nil, bosherr.WrapError(err, "Parsing URI")
	}

	switch parsed.Scheme {
	case "", "file":
		return CreateFile(f.fs, parsed.Path)
	case "http", "https":
		return CreateHTTP(http.DefaultClient, uri)
	case "s3":
		secure := true

		split := strings.SplitN(parsed.Path, "/", 3)
		if len(split) != 3 {
			return nil, fmt.Errorf("Invalid s3 bucket/object path: %s", parsed.Path)
		}

		minioEndpoint := parsed.Hostname()
		if endpointRegex.MatchString(minioEndpoint) {
			minioEndpoint = "s3.amazonaws.com"
		}

		client, err := minio.New(minioEndpoint, os.Getenv("AWS_ACCESS_KEY_ID"), os.Getenv("AWS_SECRET_ACCESS_KEY"), secure)
		if err != nil {
			return nil, bosherr.WrapError(err, "Creating s3 client")
		}

		return CreateS3(client, secure, parsed.Hostname(), split[1], split[2])
	}

	return nil, fmt.Errorf("Unknown URI scheme: %s", parsed.Scheme)
}
