package defaultloader

import (
	"os"

	"github.com/dpb587/metalink/file/url"
	fileurl "github.com/dpb587/metalink/file/url/file"
	ftpurl "github.com/dpb587/metalink/file/url/ftp"
	httpurl "github.com/dpb587/metalink/file/url/http"
	s3url "github.com/dpb587/metalink/file/url/s3"
	"github.com/dpb587/metalink/file/url/urlutil"
)

func New() url.Loader {
	file := fileurl.NewLoader()

	return url.NewMultiLoader(
		file,
		ftpurl.Loader{},
		httpurl.Loader{},
		s3url.NewLoader(s3url.Options{
			AccessKey: os.Getenv("AWS_ACCESS_KEY_ID"),
			SecretKey: os.Getenv("AWS_SECRET_ACCESS_KEY"),
			RoleARN:   os.Getenv("AWS_ROLE_ARN"),
		}),
		urlutil.NewEmptySchemeLoader(file),
	)
}
