package transfer

import (
	"github.com/cheggaaa/pb"
	"github.com/dpb587/metalink"
	"github.com/dpb587/metalink/file"
	"github.com/dpb587/metalink/verification"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . Transfer
type Transfer interface {
	TransferFile(metalink.File, file.Reference, *pb.ProgressBar, verification.VerificationResultReporter) error
}
