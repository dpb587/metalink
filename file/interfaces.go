package file

import (
	"io"

	"github.com/cheggaaa/pb"
)

//go:generate counterfeiter . Reference
type Reference interface {
	Name() (string, error)
	Size() (uint64, error)

	Reader() (io.ReadCloser, error)
	ReaderURI() string

	WriteFrom(Reference, *pb.ProgressBar) error
}
