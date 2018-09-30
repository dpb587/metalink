package defaultloader

import (
	"os"

	"github.com/dpb587/metalink/file/url"
	fileurl "github.com/dpb587/metalink/file/url/file"
	ftpurl "github.com/dpb587/metalink/file/url/ftp"
	httpurl "github.com/dpb587/metalink/file/url/http"
	s3url "github.com/dpb587/metalink/file/url/s3"
)

func New() url.Loader {
	file := fileurl.NewLoader()

	loader := url.NewLoaderFactory()
	loader.Add(file)
	loader.Add(ftpurl.Loader{})
	loader.Add(httpurl.Loader{})
	loader.Add(s3url.NewLoader(os.Getenv("AWS_ACCESS_KEY_ID"), os.Getenv("AWS_SECRET_ACCESS_KEY")))
	loader.Add(fileurl.NewEmptyLoader(file))

	return loader
}
