package blobreceipt_test

import (
	. "github.com/dpb587/blob-receipt"

	"github.com/dpb587/blob-receipt/crypto"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("BlobReceiptDigest", func() {
	var subject BlobReceiptDigest

	Describe("Algorithms", func() {
		It("enumerates algorithm keys", func() {
			subject = BlobReceiptDigest{
				"sha1": "something",
				"md16": "something-else",
			}

			algorithms := subject.Algorithms()

			Expect(algorithms).To(HaveLen(2))
			Expect(algorithms).To(ContainElement("sha1"))
			Expect(algorithms).To(ContainElement("md16"))
		})

		It("does not require an element", func() {
			subject = BlobReceiptDigest{}

			Expect(subject.Algorithms()).To(HaveLen(0))
		})
	})

	Describe("Get", func() {
		BeforeEach(func() {
			subject = BlobReceiptDigest{
				"sha1": "something",
				"md16": "something-else",
			}
		})

		It("returns digests", func() {
			digest, err := subject.Get("sha1")

			Expect(err).ToNot(HaveOccurred())
			Expect(digest).ToNot(BeNil())
			Expect(digest.Algorithm().Name()).To(Equal("sha1"))
			Expect(crypto.GetDigestHash(digest)).To(Equal("something"))
		})

		It("errors when algorithm is not recognized", func() {
			_, err := subject.Get("md16")

			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("Getting algorithm"))
		})

		It("errors when algorithm is not recognized", func() {
			_, err := subject.Get("superhash")

			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("Unknown digest: superhash"))
		})
	})
})
