package storage

import (
	"os"

	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	boshsys "github.com/cloudfoundry/bosh-utils/system"
	"github.com/dpb587/metalink"
)

type FileSystem struct {
	fs   boshsys.FileSystem
	path string
}

var _ Storage = FileSystem{}

func CreateFileSystem(fs boshsys.FileSystem, path string) (Storage, error) {
	absPath, err := fs.ExpandPath(path)
	if err != nil {
		return nil, bosherr.WrapError(err, "Expanding path")
	}

	return FileSystem{
		fs:   fs,
		path: absPath,
	}, nil
}

func (s FileSystem) String() string {
	return s.path
}

func (s FileSystem) Exists() (bool, error) {
	return s.fs.FileExists(s.path), nil
}

func (s FileSystem) Get() (metalink.Metalink, error) {
	file, err := s.fs.OpenFile(s.path, os.O_RDONLY, 0)
	if err != nil {
		return metalink.Metalink{}, bosherr.WrapError(err, "Opening file for writing")
	}

	return ReadMetalink(file)
}

func (s FileSystem) Put(receipt metalink.Metalink) error {
	file, err := s.fs.OpenFile(s.path, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		return bosherr.WrapError(err, "Opening file for writing")
	}

	err = WriteMetalink(file, receipt)
	if err != nil {
		return bosherr.WrapError(err, "Writing receipt to file")
	}

	err = file.Close()
	if err != nil {
		return bosherr.WrapError(err, "Closing file")
	}

	return nil
}
