package args_test

import (
	"fmt"
	"time"

	. "github.com/dpb587/blob-receipt/cli/args"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Time", func() {
	var subject Time

	Describe("UnmarshalFlag", func() {
		BeforeEach(func() {
			subject = Time{}
		})

		It("parses RFC3339 formats", func() {
			err := subject.UnmarshalFlag("2016-03-02T01:12:11Z")

			Expect(err).ToNot(HaveOccurred())
			Expect(subject.Unix()).To(Equal(int64(1456881131)))
		})

		It("errors when = is missing", func() {
			err := subject.UnmarshalFlag("value1")

			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("Parsing time argument"))
		})
	})

	Describe("String", func() {
		It("stringifies", func() {
			parsed, _ := time.Parse(time.RFC3339, "2016-03-02T01:12:11Z")
			subject = Time{
				Time: parsed,
			}

			Expect(fmt.Sprintf("%s", subject)).To(Equal("2016-03-02T01:12:11Z"))
		})
	})
})
