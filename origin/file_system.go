package origin

import (
	"io"
	"os"
	"path/filepath"
	"time"

	boshcry "github.com/cloudfoundry/bosh-utils/crypto"
	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	boshsys "github.com/cloudfoundry/bosh-utils/system"
)

type FileSystem struct {
	fs   boshsys.FileSystem
	path string
}

var _ Origin = FileSystem{}

func CreateFileSystem(fs boshsys.FileSystem, path string) (Origin, error) {
	absPath, err := fs.ExpandPath(path)
	if err != nil {
		return nil, bosherr.WrapError(err, "Expanding path")
	}

	return FileSystem{
		fs:   fs,
		path: absPath,
	}, nil
}

func (o FileSystem) String() string {
	return o.path
}

func (o FileSystem) Digest(algorithm boshcry.Algorithm) (boshcry.Digest, error) {
	reader, err := o.Reader()
	if err != nil {
		return nil, bosherr.WrapError(err, "Digesting file")
	}

	return algorithm.CreateDigest(reader)
}

func (o FileSystem) Name() (string, error) {
	return filepath.Base(o.path), nil
}

func (o FileSystem) Size() (uint64, error) {
	stat, err := o.fs.Stat(o.path)
	if err != nil {
		return 0, bosherr.WrapError(err, "Checking file size")
	}

	return uint64(stat.Size()), nil
}

func (o FileSystem) Time() (time.Time, error) {
	stat, err := o.fs.Stat(o.path)
	if err != nil {
		return time.Time{}, bosherr.WrapError(err, "Checking file time")
	}

	return stat.ModTime(), nil
}

func (o FileSystem) Reader() (io.Reader, error) {
	reader, err := o.fs.OpenFile(o.path, os.O_RDONLY, 0000)
	if err != nil {
		return nil, bosherr.WrapError(err, "Opening file")
	}

	return reader, nil
}
