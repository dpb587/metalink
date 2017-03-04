package storage

import (
	boshsys "github.com/cloudfoundry/bosh-utils/system"
)

type defaultFactory struct {
	fs boshsys.FileSystem
}

var _ StorageFactory = defaultFactory{}

func NewDefaultFactory(fs boshsys.FileSystem) StorageFactory {
	return defaultFactory{
		fs: fs,
	}
}

func (f defaultFactory) New(path string) (Storage, error) {
	return CreateFileSystem(f.fs, path)
}
