package git

import (
	"fmt"
	"net/url"
	"os"
	"regexp"
	"strings"

	"github.com/dpb587/metalink/repository/source"
	minio "github.com/minio/minio-go"

	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	boshsys "github.com/cloudfoundry/bosh-utils/system"
)

// http://docs.aws.amazon.com/general/latest/gr/rande.html#s3_region
var endpointRegex = regexp.MustCompile(`^s3(\.(dualstack\.)?|\-)[^\.]+\.amazonaws.com$`)

type Factory struct{}

var _ source.Factory = &Factory{}

func NewFactory(fs boshsys.FileSystem, cmdRunner boshsys.CmdRunner) Factory {
	return Factory{}
}

func (f Factory) Schemes() []string {
	return []string{"s3"}
}

func (f Factory) Create(uri string) (source.Source, error) {
	parsed, err := url.Parse(uri)
	if err != nil {
		return nil, bosherr.WrapError(err, "Parsing URI")
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

	client, err := minio.New(minioEndpoint, os.Getenv("AWS_ACCESS_KEY_ID"), os.Getenv("AWS_SECRET_ACCESS_KEY"), secure)
	if err != nil {
		return nil, bosherr.WrapError(err, "Creating s3 client")
	}

	return NewSource(uri, client, secure, split[1], split[2]), nil
}
