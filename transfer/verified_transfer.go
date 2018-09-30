package transfer

import (
	"github.com/cheggaaa/pb"
	"github.com/dpb587/metalink"
	"github.com/dpb587/metalink/file"
	"github.com/dpb587/metalink/file/metaurl"
	"github.com/dpb587/metalink/file/url"
	"github.com/dpb587/metalink/verification"
	"github.com/pkg/errors"
)

type VerifiedTransfer struct {
	metaurlLoader metaurl.Loader
	urlLoader     url.Loader
	verifier      verification.Verifier
}

func NewVerifiedTransfer(metaurlLoader metaurl.Loader, urlLoader url.Loader, verifier verification.Verifier) Transfer {
	return VerifiedTransfer{
		metaurlLoader: metaurlLoader,
		urlLoader:     urlLoader,
		verifier:      verifier,
	}
}

func (t VerifiedTransfer) TransferFile(meta4file metalink.File, local file.Reference, progress *pb.ProgressBar, verificationResultReporter verification.VerificationResultReporter) error {
	sources := newSourceList(meta4file.MetaURLs, meta4file.URLs)

	errs := []error{}

	for _, source := range sources {
		var err error

		if source.URL != nil {
			err = t.transferFileURL(meta4file, local, progress, verificationResultReporter, *source.URL)
		} else if source.MetaURL != nil {
			err = t.transferFileMetaURL(meta4file, local, progress, verificationResultReporter, *source.MetaURL)
		} else {
			panic("missing url or metaurl")
		}

		if err == nil {
			return nil
		}

		errs = append(errs, err)
	}

	if len(errs) == 0 {
		return errors.New("no valid url found")
	}

	progress.Finish()

	return errors.New(errs[0].Error())
}

func (t VerifiedTransfer) transferFileURL(meta4file metalink.File, local file.Reference, progress *pb.ProgressBar, verificationResultReporter verification.VerificationResultReporter, source metalink.URL) error {
	remote, err := t.urlLoader.Load(source)
	if err != nil {
		return errors.Wrap(err, "Parsing source file")
	}

	progress.Start()

	err = local.WriteFrom(remote, progress)
	if err != nil {
		return errors.Wrap(err, "Transferring file")
	}

	progress.Finish()

	result := t.verifier.Verify(local, meta4file)

	if verificationResultReporter != nil {
		verificationResultReporter.ReportVerificationResult(result)
	}

	if result.Error() != nil {
		return errors.Wrap(result.Error(), "Verifying file")
	}

	return nil
}

func (t VerifiedTransfer) transferFileMetaURL(meta4file metalink.File, local file.Reference, progress *pb.ProgressBar, _ verification.VerificationResultReporter, source metalink.MetaURL) error {
	remote, err := t.metaurlLoader.Load(source)
	if err != nil {
		return errors.Wrap(err, "Parsing source file")
	}

	progress.Start()

	err = local.WriteFrom(remote, progress)
	if err != nil {
		return errors.Wrap(err, "Transferring file")
	}

	// @todo; metaurl downloads may not match
	// err = t.verifier.Verify(local, meta4file)
	// if err != nil {
	// 	return errors.Wrap(err, "Verifying file")
	// }

	progress.Finish()

	return nil
}
