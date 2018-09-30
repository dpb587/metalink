package file

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/cheggaaa/pb"
	"github.com/pkg/errors"
	"github.com/dpb587/metalink/file"
)

type Reference struct {
	path string
}

var _ file.Reference = Reference{}

func NewReference(path string) Reference {
	return Reference{
		path: path,
	}
}

func (o Reference) Name() (string, error) {
	return filepath.Base(o.path), nil
}

func (o Reference) Size() (uint64, error) {
	stat, err := os.Stat(o.path)
	if err != nil {
		return 0, errors.Wrap(err, "Checking file size")
	}

	return uint64(stat.Size()), nil
}

func (o Reference) Reader() (io.ReadCloser, error) {
	reader, err := os.OpenFile(o.path, os.O_RDONLY, 0000)
	if err != nil {
		return nil, errors.Wrap(err, "Opening file for reading")
	}

	return reader, nil
}

func (o Reference) ReaderURI() string {
	return fmt.Sprintf("file://%s", o.path)
}

func (o Reference) WriteFrom(from file.Reference, progress *pb.ProgressBar) error {
	reader, err := from.Reader()
	if err != nil {
		return errors.Wrap(err, "Opening from")
	}

	defer reader.Close()

	writer, err := os.OpenFile(o.path, os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		return errors.Wrap(err, "Opening file for writing")
	}

	defer writer.Close()

	_, err = io.Copy(writer, progress.NewProxyReader(reader))

	return err
}
