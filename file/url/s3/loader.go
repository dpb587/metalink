package s3

import (
	"fmt"
	neturl "net/url"
	"os"
	"regexp"
	"strings"

	"github.com/pkg/errors"
	"github.com/dpb587/metalink"
	"github.com/dpb587/metalink/file"
	"github.com/dpb587/metalink/file/url"
	minio "github.com/minio/minio-go"
)

// http://docs.aws.amazon.com/general/latest/gr/rande.html#s3_region
var endpointRegex = regexp.MustCompile(`^s3(\.(dualstack\.)?|\-)[^\.]+\.amazonaws.com$`)

type Loader struct{}

var _ url.Loader = &Loader{}

func (f Loader) Schemes() []string {
	return []string{
		"s3",
	}
}

func (f Loader) Load(source metalink.URL) (file.Reference, error) {
	parsed, err := neturl.Parse(source.URL)
	if err != nil {
		return nil, errors.Wrap(err, "Parsing URI")
	}

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
		return nil, errors.Wrap(err, "Creating s3 client")
	}

	return NewReference(client, secure, parsed.Hostname(), split[1], split[2]), nil
}
