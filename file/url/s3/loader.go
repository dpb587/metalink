package s3

import (
	"fmt"
	neturl "net/url"
	"regexp"
	"strings"

	"github.com/dpb587/metalink"
	"github.com/dpb587/metalink/file"
	"github.com/dpb587/metalink/file/url"
	minio "github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/pkg/errors"
)

// http://docs.aws.amazon.com/general/latest/gr/rande.html#s3_region
var endpointRegex = regexp.MustCompile(`^s3(\.(dualstack\.)?|\-)[^\.]+\.amazonaws.com$`)

type loader struct {
	options Options
}

var _ url.Loader = &loader{}

func NewLoader(options Options) url.Loader {
	return &loader{options}
}

func (f loader) SupportsURL(source metalink.URL) bool {
	parsed, err := neturl.Parse(source.URL)
	if err != nil {
		return false
	}

	if parsed.Scheme == "s3" {
		return true
	}

	if endpointRegex.MatchString(parsed.Hostname()) {
		return true
	}

	return false
}

func (f loader) LoadURL(source metalink.URL) (file.Reference, error) {
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

	var minioCreds *credentials.Credentials

	if f.options.RoleARN == "" {
		minioCreds = credentials.NewStaticV4(f.options.AccessKey, f.options.SecretKey, "")
	} else {
		if minioEndpoint != "s3.amazonaws.com" {
			return nil, errors.New("Role ARN is only supported for S3 endpoints")
		}

		minioCreds, err = credentials.NewSTSAssumeRole(
			"https://sts.amazonaws.com",
			credentials.STSAssumeRoleOptions{
				AccessKey:       f.options.AccessKey,
				SecretKey:       f.options.SecretKey,
				Location:        "us-east-1",
				RoleARN:         f.options.RoleARN,
				RoleSessionName: "metalink-session",
			},
		)
		if err != nil {
			return nil, errors.Wrap(err, "Unable to authenticate as assumed role")
		}
	}

	minioOptions := &minio.Options{
		Creds:  minioCreds,
		Secure: true,
	}

	client, err := minio.New(minioEndpoint, minioOptions)
	if err != nil {
		return nil, errors.Wrap(err, "Creating s3 client")
	}

	return NewReference(client, secure, parsed.Hostname(), split[1], split[2]), nil
}
