package blobreceipt_test

import (
	"bytes"
	"errors"
	"time"

	boshcry "github.com/cloudfoundry/bosh-utils/crypto"
	. "github.com/dpb587/blob-receipt"
	"github.com/dpb587/blob-receipt/origin/originfakes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("BlobReceipt", func() {
	var subject BlobReceipt

	BeforeEach(func() {
		subject = BlobReceipt{}
	})

	Describe("SetMetadata", func() {
		It("sets new metadata", func() {
			subject.SetMetadata("key1", "value1")

			Expect(subject.Metadata).To(HaveLen(1))

			metadata := subject.Metadata[0]
			Expect(metadata.Key).To(Equal("key1"))
			Expect(metadata.Value).To(Equal("value1"))
		})

		It("appends new metadata", func() {
			subject.SetMetadata("key1", "value1")
			subject.SetMetadata("key2", "value2")

			Expect(subject.Metadata).To(HaveLen(2))

			metadata := subject.Metadata[0]
			Expect(metadata.Key).To(Equal("key1"))
			Expect(metadata.Value).To(Equal("value1"))

			metadata = subject.Metadata[1]
			Expect(metadata.Key).To(Equal("key2"))
			Expect(metadata.Value).To(Equal("value2"))
		})

		It("updates existing metadata", func() {
			subject.SetMetadata("key1", "value1")
			subject.SetMetadata("key1", "value1b")

			Expect(subject.Metadata).To(HaveLen(1))

			metadata := subject.Metadata[0]
			Expect(metadata.Key).To(Equal("key1"))
			Expect(metadata.Value).To(Equal("value1b"))
		})
	})

	Describe("SetOrigin", func() {
		It("sets new origin", func() {
			origin := BlobReceiptOrigin{"uri": "test"}
			subject.SetOrigin(origin)

			Expect(subject.Origin).To(HaveLen(1))
			Expect(subject.Origin[0]).To(Equal(origin))
		})

		It("appends new origin", func() {
			origin1 := BlobReceiptOrigin{"uri": "test1"}
			subject.SetOrigin(origin1)

			origin2 := BlobReceiptOrigin{"uri": "test2"}
			subject.SetOrigin(origin2)

			Expect(subject.Origin).To(HaveLen(2))
			Expect(subject.Origin[0]).To(Equal(origin1))
			Expect(subject.Origin[1]).To(Equal(origin2))
		})

		It("updates existing origin", func() {
			subject.SetOrigin(BlobReceiptOrigin{"uri": "test1"})

			origin2 := BlobReceiptOrigin{"uri": "test1", "custom": "two"}
			subject.SetOrigin(origin2)

			Expect(subject.Origin).To(HaveLen(1))
			Expect(subject.Origin[0]).To(Equal(origin2))
		})
	})

	Describe("Write", func() {
		It("basically works", func() {
			subject.SetMetadata("key2", "value2")
			subject.SetMetadata("key1", "value1")

			subject.Digest = BlobReceiptDigest{}
			subject.Digest["md5"] = "fake-md5-hash"
			subject.Digest["superhash"] = "fake-super-hash"

			subject.SetOrigin(BlobReceiptOrigin{"uri": "custom", "blocks": []string{"one", "two"}})
			subject.SetOrigin(BlobReceiptOrigin{"uri": "alpha"})

			then, _ := time.Parse(time.RFC3339, "2017-04-01T02:03:04")
			subject.Time = then

			subject.Name = "test.tar.gz"
			subject.Size = 123456

			buf := bytes.NewBufferString("")

			err := subject.Write(buf)

			Expect(err).ToNot(HaveOccurred())

			Expect(buf.String()).To(Equal(`{
  "digest": {
    "md5": "fake-md5-hash",
    "superhash": "fake-super-hash"
  },
  "metadata": [
    {
      "key": "key1",
      "value": "value1"
    },
    {
      "key": "key2",
      "value": "value2"
    }
  ],
  "name": "test.tar.gz",
  "origin": [
    {
      "uri": "alpha"
    },
    {
      "blocks": [
        "one",
        "two"
      ],
      "uri": "custom"
    }
  ],
  "size": 123456,
  "time": "0001-01-01T00:00:00Z"
}
`))
		})
	})

	Describe("UpdateFromOrigin", func() {
		var origin originfakes.FakeOrigin

		BeforeEach(func() {
			origin = originfakes.FakeOrigin{}
		})

		Context("Name", func() {
			BeforeEach(func() {
				subject.Name = "test.tar.gz"
			})

			It("updates name", func() {
				origin.NameReturns("other.tar.gz", nil)

				err := subject.UpdateFromOrigin(&origin, nil)

				Expect(err).ToNot(HaveOccurred())
				Expect(subject.Name).To(Equal("other.tar.gz"))
			})

			It("propagates errors", func() {
				origin.NameReturns("", errors.New("fake-error"))

				err := subject.UpdateFromOrigin(&origin, nil)

				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("Sourcing blob name"))
				Expect(err.Error()).To(ContainSubstring("fake-error"))
			})
		})

		Context("Size", func() {
			BeforeEach(func() {
				subject.Size = 12345
			})

			It("updates size", func() {
				origin.SizeReturns(98765, nil)

				err := subject.UpdateFromOrigin(&origin, nil)

				Expect(err).ToNot(HaveOccurred())
				Expect(subject.Size).To(Equal(uint64(98765)))
			})

			It("propagates errors", func() {
				origin.SizeReturns(0, errors.New("fake-error"))

				err := subject.UpdateFromOrigin(&origin, nil)

				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("Sourcing blob size"))
				Expect(err.Error()).To(ContainSubstring("fake-error"))
			})
		})

		Context("Time", func() {
			BeforeEach(func() {
				then, _ := time.Parse(time.RFC3339, "2017-04-01T02:03:04Z")

				subject.Time = then
			})

			It("updates time", func() {
				then, _ := time.Parse(time.RFC3339, "2011-08-02T04:06:08Z")
				origin.TimeReturns(then, nil)

				err := subject.UpdateFromOrigin(&origin, nil)

				Expect(err).ToNot(HaveOccurred())
				Expect(subject.Time).To(Equal(then))
			})

			It("propagates errors", func() {
				origin.TimeReturns(time.Time{}, errors.New("fake-error"))

				err := subject.UpdateFromOrigin(&origin, nil)

				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("Sourcing blob time"))
				Expect(err.Error()).To(ContainSubstring("fake-error"))
			})
		})

		Context("Digests", func() {
			It("handles known algorithms", func() {
				origin.DigestReturns(boshcry.NewDigest(boshcry.DigestAlgorithmSHA1, "fake-digest"), nil)

				err := subject.UpdateFromOrigin(&origin, []string{"sha1"})

				Expect(err).ToNot(HaveOccurred())
				Expect(subject.Digest).To(HaveLen(1))
				Expect(subject.Digest["sha1"]).To(Equal("fake-digest"))
			})

			It("errors with unknown algorithms", func() {
				err := subject.UpdateFromOrigin(&origin, []string{"superhash"})

				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("Loading digest algorithm"))
			})

			It("errors with unknown algorithms", func() {
				origin.DigestReturns(nil, errors.New("fake-error"))

				err := subject.UpdateFromOrigin(&origin, []string{"sha1"})

				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("Sourcing blob sha1 digest"))
			})
		})
	})
})
