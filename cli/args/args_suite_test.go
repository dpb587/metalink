package args_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestArgs(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "github.com/dpb587/blob-receipt/cli/args")
}
