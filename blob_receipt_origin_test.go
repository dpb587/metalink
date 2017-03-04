package blobreceipt_test

import (
	. "github.com/dpb587/blob-receipt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("BlobReceiptOrigin", func() {
	var subject BlobReceiptOrigin

	Describe("URI", func() {
		It("returns the uri key", func() {
			subject = BlobReceiptOrigin{
				"uri": "somewhere",
			}

			Expect(subject.URI()).To(Equal("somewhere"))
		})
	})
})
