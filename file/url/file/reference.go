package file

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/cheggaaa/pb"
	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	boshsys "github.com/cloudfoundry/bosh-utils/system"
	"github.com/dpb587/metalink/file"
)

type Reference struct {
	fs   boshsys.FileSystem
	path string
}

var _ file.Reference = Reference{}

func NewReference(fs boshsys.FileSystem, path string) Reference {
	return Reference{
		fs:   fs,
		path: path,
	}
}

func (o Reference) Name() (string, error) {
	return filepath.Base(o.path), nil
}

func (o Reference) Size() (uint64, error) {
	stat, err := o.fs.Stat(o.path)
	if err != nil {
		return 0, bosherr.WrapError(err, "Checking file size")
	}

	return uint64(stat.Size()), nil
}

func (o Reference) Reader() (io.ReadCloser, error) {
	reader, err := o.fs.OpenFile(o.path, os.O_RDONLY, 0000)
	if err != nil {
		return nil, bosherr.WrapError(err, "Opening file for reading")
	}

	return reader, nil
}

func (o Reference) ReaderURI() string {
	return fmt.Sprintf("file://%s", o.path)
}

func (o Reference) WriteFrom(from file.Reference, progress *pb.ProgressBar) error {
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
