package crypto_test

import (
	"bytes"
	"io"
	"testing/iotest"

	. "github.com/dpb587/metalink/crypto"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Md5", func() {
	Describe("Name", func() {
		It("knows its hash", func() {
			Expect(DigestAlgorithmMD5.Name()).To(Equal("md5"))
		})
	})

	Describe("CreateDigest", func() {
		var reader io.Reader

		BeforeEach(func() {
			reader = bytes.NewReader([]byte("something different"))
		})

		It("calculates the hash", func() {
			digest, err := DigestAlgorithmMD5.CreateDigest(reader)

			Expect(err).ToNot(HaveOccurred())
			Expect(digest).ToNot(BeNil())
			Expect(digest.String()).To(Equal("md5:0c895270853c3023cf2741bdfdab14e2"))
		})

		It("propagates read errors", func() {
			reader = iotest.TimeoutReader(reader)

			_, err := DigestAlgorithmMD5.CreateDigest(reader)

			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("Reading file for digest"))
		})
	})
})
