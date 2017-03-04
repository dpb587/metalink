package args_test

import (
	. "github.com/dpb587/blob-receipt/cli/args"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Origin", func() {
	var subject Origin

	Describe("UnmarshalFlag", func() {
		BeforeEach(func() {
			subject = Origin{}
		})

		It("handles plain string URIs", func() {
			err := subject.UnmarshalFlag("https://example.com/somewhere")

			Expect(err).ToNot(HaveOccurred())
			Expect(subject.URI()).To(Equal("https://example.com/somewhere"))
		})

		It("handles plain yaml URIs", func() {
			err := subject.UnmarshalFlag(`{"uri":"https://example.com/somewhere-else"}`)

			Expect(err).ToNot(HaveOccurred())
			Expect(subject.URI()).To(Equal("https://example.com/somewhere-else"))
		})

		It("handles extra yaml properties", func() {
			err := subject.UnmarshalFlag(`{"uri":"https://example.com/somewhere-else","blocks":["one","two"]}`)

			Expect(err).ToNot(HaveOccurred())
			Expect(subject.URI()).To(Equal("https://example.com/somewhere-else"))
			Expect(subject.BlobReceiptOrigin["blocks"]).To(Equal([]interface{}{"one", "two"}))
		})

		It("propagates yaml parsing errors", func() {
			err := subject.UnmarshalFlag(`{"uri":"trailing`)

			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("Unmarshaling JSON origin"))
		})
	})
})
