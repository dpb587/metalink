package storage

import (
	"io/ioutil"
	"os"

	"github.com/pkg/errors"
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
		return nil, errors.Wrap(err, "Expanding path")
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
		return metalink.Metalink{}, errors.Wrap(err, "Opening file for reading")
	}

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return metalink.Metalink{}, errors.Wrap(err, "Reading XML")
	}

	meta4 := metalink.Metalink{}

	err = metalink.Unmarshal(bytes, &meta4)
	if err != nil {
		return metalink.Metalink{}, errors.Wrap(err, "Unmarshaling")
	}

	return meta4, nil
}

func (s FileSystem) Put(receipt metalink.Metalink) error {
	bytes, err := metalink.Marshal(receipt)
	if err != nil {
		return errors.Wrap(err, "Marshaling")
	}

	file, err := s.fs.OpenFile(s.path, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		return errors.Wrap(err, "Opening file for writing")
	}

	file.Write([]byte(`<?xml version="1.0" encoding="utf-8"?>`))
	file.Write([]byte("\n"))
	file.Write(bytes)

	err = file.Close()
	if err != nil {
		return errors.Wrap(err, "Closing file")
	}

	return nil
}
