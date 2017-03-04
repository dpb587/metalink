package storage

import (
	"encoding/json"
	"os"

	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	boshsys "github.com/cloudfoundry/bosh-utils/system"
	blobreceipt "github.com/dpb587/blob-receipt"
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

// @todo path or absolute path
func (s FileSystem) String() string {
	return s.path
}

func (s FileSystem) Exists() (bool, error) {
	return s.fs.FileExists(s.path), nil
}

func (s FileSystem) Get() (blobreceipt.BlobReceipt, error) {
	receipt := blobreceipt.BlobReceipt{}

	bytes, err := s.fs.ReadFile(s.path)
	if err != nil {
		return receipt, bosherr.WrapError(err, "Reading file")
	}

	err = json.Unmarshal(bytes, &receipt)
	if err != nil {
		return receipt, bosherr.WrapError(err, "Unmarshaling JSON")
	}

	return receipt, nil
}

func (s FileSystem) Put(receipt blobreceipt.BlobReceipt) error {
	file, err := s.fs.OpenFile(s.path, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return bosherr.WrapError(err, "Opening file for writing")
	}

	err = receipt.Write(file)
	if err != nil {
		return bosherr.WrapError(err, "Writing receipt to file")
	}

	err = file.Close()
	if err != nil {
		return bosherr.WrapError(err, "Closing file")
	}

	return nil
}
