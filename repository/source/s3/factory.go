package git

import (
	"fmt"
	"net/url"
	"os"
	"regexp"
	"strings"

	"github.com/dpb587/metalink/repository/source"
	minio "github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/pkg/errors"
)

// http://docs.aws.amazon.com/general/latest/gr/rande.html#s3_region
var endpointRegex = regexp.MustCompile(`^s3(\.(dualstack\.)?|\-)[^\.]+\.amazonaws.com$`)

type Factory struct{}

var _ source.Factory = &Factory{}

func NewFactory() Factory {
	return Factory{}
}

func (f Factory) Schemes() []string {
	return []string{"s3"}
}

func (f Factory) Create(uri string, options map[string]interface{}) (source.Source, error) {
	parsed, err := url.Parse(uri)
	if err != nil {
		return nil, errors.Wrap(err, "Parsing URI")
	}

	secure := true

	split := strings.SplitN(parsed.Path, "/", 3)
	if len(split) != 3 {
		return nil, fmt.Errorf("Invalid s3 bucket/prefix path: %s", parsed.Path)
	}

	minioEndpoint := parsed.Hostname()
	if endpointRegex.MatchString(minioEndpoint) {
		minioEndpoint = "s3.amazonaws.com"
	}

	if parsed.Port() != "" && parsed.Port() != "443" {
		minioEndpoint = fmt.Sprintf("%s:%s", minioEndpoint, parsed.Port())
	}

	var accessKey, secretKey string
	var optionValid bool

	accessKey = os.Getenv("AWS_ACCESS_KEY_ID")
	secretKey = os.Getenv("AWS_SECRET_ACCESS_KEY")

	if accessKeyOpt, found := options["access_key"]; found {
		if accessKey, optionValid = accessKeyOpt.(string); !optionValid {
			return nil, errors.New("Option 'access_key' must be string")
		}
	}

	if secretKeyOpt, found := options["secret_key"]; found {
		if secretKey, optionValid = secretKeyOpt.(string); !optionValid {
			return nil, errors.New("Option 'secret_key' must be string")
		}
	}

	if parsed.User != nil {
		accessKey = parsed.User.Username()
		secretKey, _ = parsed.User.Password()
	}

	minioCreds := credentials.NewStaticV4(accessKey, secretKey, "")

	minioOptions := &minio.Options{
		Creds:  minioCreds,
		Secure: true,
	}

	client, err := minio.New(minioEndpoint, minioOptions)
	if err != nil {
		return nil, errors.Wrap(err, "Creating s3 client")
	}

	return NewSource(uri, client, secure, split[1], split[2]), nil
}
