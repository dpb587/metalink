package verification

import (
	"fmt"
	"io"
)

type prefixedVerificationResultReporter struct {
	writer io.Writer
	prefix string
}

func NewPrefixedVerificationResultReporter(writer io.Writer, prefix string) VerificationResultReporter {
	return &prefixedVerificationResultReporter{
		writer: writer,
		prefix: prefix,
	}
}

func (cr prefixedVerificationResultReporter) ReportVerificationResult(result VerificationResult) error {
	if multi, ok := result.(MultipleVerificationResults); ok {
		for _, result := range multi.VerificationResults() {
			err := cr.ReportVerificationResult(result)
			if err != nil {
				return err
			}
		}

		return nil
	}

	var err error

	if result.Error() != nil {
		_, err = fmt.Fprintf(cr.writer, "%sERROR: %s: %s\n", cr.prefix, result.Verifier(), result.Confirmation())
	} else {
		_, err = fmt.Fprintf(cr.writer, "%s%s: %s\n", cr.prefix, result.Verifier(), result.Confirmation())
	}

	return err
}
