package transfer

import (
	"github.com/cheggaaa/pb"
	"github.com/dpb587/metalink"
	"github.com/dpb587/metalink/file"
	"github.com/dpb587/metalink/verification"
)

//go:generate counterfeiter . Transfer
type Transfer interface {
	TransferFile(metalink.File, file.Reference, *pb.ProgressBar, verification.VerificationResultReporter) error
}
