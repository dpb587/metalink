package origin

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/cheggaaa/pb"
	boshcry "github.com/cloudfoundry/bosh-utils/crypto"
	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	boshsys "github.com/cloudfoundry/bosh-utils/system"
)

type File struct {
	fs   boshsys.FileSystem
	path string
}

var _ Origin = File{}

func CreateFile(fs boshsys.FileSystem, path string) (Origin, error) {
	absPath, err := fs.ExpandPath(path)
	if err != nil {
		return nil, bosherr.WrapError(err, "Expanding path")
	}

	return File{
		fs:   fs,
		path: absPath,
	}, nil
}

func (o File) Digest(algorithm boshcry.Algorithm) (boshcry.Digest, error) {
	reader, err := o.Reader()
	if err != nil {
		return nil, bosherr.WrapError(err, "Digesting file")
	}

	return algorithm.CreateDigest(reader)
}

func (o File) Name() (string, error) {
	return filepath.Base(o.path), nil
}

func (o File) Size() (uint64, error) {
	stat, err := o.fs.Stat(o.path)
	if err != nil {
		return 0, bosherr.WrapError(err, "Checking file size")
	}

	return uint64(stat.Size()), nil
}

func (o File) Reader() (io.ReadCloser, error) {
	reader, err := o.fs.OpenFile(o.path, os.O_RDONLY, 0000)
	if err != nil {
		return nil, bosherr.WrapError(err, "Opening file for reading")
	}

	return reader, nil
}

func (o File) ReaderURI() string {
	return fmt.Sprintf("file://%s", o.path)
}

func (o File) WriteFrom(from Origin, progress *pb.ProgressBar) error {
	reader, err := from.Reader()
	if err != nil {
		return bosherr.WrapError(err, "Opening from")
	}

	defer reader.Close()

	writer, err := o.fs.OpenFile(o.path, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return bosherr.WrapError(err, "Opening file for writing")
	}

	defer writer.Close()

	_, err = io.Copy(writer, progress.NewProxyReader(reader))

	return err
}
