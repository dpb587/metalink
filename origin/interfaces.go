package origin

import (
	"io"
	"time"

	boshcry "github.com/cloudfoundry/bosh-utils/crypto"
)

//go:generate counterfeiter . Origin
type Origin interface {
	String() string

	Digest(boshcry.Algorithm) (boshcry.Digest, error)
	Name() (string, error)
	Size() (uint64, error)
	Time() (time.Time, error)

	Reader() (io.Reader, error)
}

//go:generate counterfeiter . OriginFactory
type OriginFactory interface {
	New(string) (Origin, error)
}
