package origin_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestOrigin(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "github.com/dpb587/blob-receipt/origin")
}
