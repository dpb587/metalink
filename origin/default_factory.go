package origin

import (
	boshsys "github.com/cloudfoundry/bosh-utils/system"
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

func (f defaultFactory) New(path string) (Origin, error) {
	return CreateFileSystem(f.fs, path)
}
