package args_test

import (
	. "github.com/dpb587/blob-receipt/cli/args"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("DigestAlgorithm", func() {
	var subject DigestAlgorithm

	BeforeEach(func() {
		subject = DigestAlgorithm{}
	})

	Describe("UnmarshalFlag", func() {
		It("knows sha1", func() {
			err := subject.UnmarshalFlag("sha1")

			Expect(err).ToNot(HaveOccurred())
			Expect(subject.Name()).To(Equal("sha1"))
		})

		It("knows sha256", func() {
			err := subject.UnmarshalFlag("sha256")

			Expect(err).ToNot(HaveOccurred())
			Expect(subject.Name()).To(Equal("sha256"))
		})

		It("knows sha512", func() {
			err := subject.UnmarshalFlag("sha512")

			Expect(err).ToNot(HaveOccurred())
			Expect(subject.Name()).To(Equal("sha512"))
		})

		It("errors with unrecognized algorithm", func() {
			err := subject.UnmarshalFlag("unknown1")

			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("Unsupported digest algorithm: unknown1"))
		})
	})
})
