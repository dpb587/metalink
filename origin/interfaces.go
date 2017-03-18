package origin

import (
	"io"
	"time"

	"github.com/cheggaaa/pb"
	boshcry "github.com/cloudfoundry/bosh-utils/crypto"
)

//go:generate counterfeiter . Origin
type Origin interface {
	Digest(boshcry.Algorithm) (boshcry.Digest, error)
	Name() (string, error)
	Size() (uint64, error)
	Time() (time.Time, error)

	Reader() (io.ReadCloser, error)
	ReaderURI() string

	WriteFrom(Origin, *pb.ProgressBar) error
}

//go:generate counterfeiter . OriginFactory
type OriginFactory interface {
	New(string) (Origin, error)
}
