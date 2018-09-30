package verification

import (
	"fmt"
	"io"

	"github.com/dpb587/metalink"
)

type simpleVerificationResultReporter struct {
	writer io.Writer
}

func NewSimpleVerificationResultReporter(writer io.Writer) VerificationResultReporter {
	return &simpleVerificationResultReporter{
		writer: writer,
	}
}

func (cr simpleVerificationResultReporter) ReportVerificationResult(file metalink.File, result VerificationResult) error {
	if multi, ok := result.(MultipleVerificationResults); ok {
		for _, result := range multi.VerificationResults() {
			err := cr.ReportVerificationResult(file, result)
			if err != nil {
				return err
			}
		}

		return nil
	}

	var err error

	if result.Error() != nil {
		_, err = fmt.Fprintf(cr.writer, "%s: %s: INVALID: %s\n", file.Name, result.Verifier(), result.Confirmation())
	} else {
		_, err = fmt.Fprintf(cr.writer, "%s: %s: %s\n", file.Name, result.Verifier(), result.Confirmation())
	}

	return err
}
