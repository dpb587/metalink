package signature

import "io"

type Signature interface {
	DefaultTrustStore() (io.Reader, error)
	Verify(io.Reader, io.Reader, io.Reader) (string, error)
}
