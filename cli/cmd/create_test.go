package cmd_test

import (
	"errors"
	"time"

	blobreceipt "github.com/dpb587/blob-receipt"
	"github.com/dpb587/blob-receipt/cli/args"
	. "github.com/dpb587/blob-receipt/cli/cmd"

	"github.com/dpb587/blob-receipt/origin/originfakes"
	"github.com/dpb587/blob-receipt/storage/storagefakes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Create", func() {
	var subject Create
	var originFactory originfakes.FakeOriginFactory
	var storageFactory storagefakes.FakeStorageFactory

	BeforeEach(func() {
		originFactory = originfakes.FakeOriginFactory{}
		storageFactory = storagefakes.FakeStorageFactory{}

		subject = Create{
			Args:           CreateArgs{},
			OriginFactory:  &originFactory,
			StorageFactory: &storageFactory,
		}
	})

	It("handles storage factory errors", func() {
		storageFactory.NewReturns(nil, errors.New("fake-err"))

		err := subject.Execute([]string{})

		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(ContainSubstring("Loading storage"))
		Expect(err.Error()).To(ContainSubstring("fake-err"))
	})

	Context("with receipt storage", func() {
		var storage storagefakes.FakeStorage

		BeforeEach(func() {
			storage = storagefakes.FakeStorage{}
			storageFactory.NewReturns(&storage, nil)
		})

		It("handles receipt exists errors", func() {
			storage.ExistsReturns(false, errors.New("fake-err"))

			err := subject.Execute([]string{})

			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("Checking if receipt exists"))
			Expect(err.Error()).To(ContainSubstring("fake-err"))
		})

		It("sets name", func() {
			v := "testme"
			subject.Name = &v

			err := subject.Execute([]string{})

			Expect(err).ToNot(HaveOccurred())
			Expect(storage.PutCallCount()).To(Equal(1))

			put := storage.PutArgsForCall(0)
			Expect(put.Name).To(Equal("testme"))
		})

		It("sets time", func() {
			v := args.MustNewTime("2017-02-03T04:05:06Z")
			subject.Time = &v

			err := subject.Execute([]string{})

			Expect(err).ToNot(HaveOccurred())
			Expect(storage.PutCallCount()).To(Equal(1))

			put := storage.PutArgsForCall(0)
			Expect(put.Time.Format(time.RFC3339)).To(Equal("2017-02-03T04:05:06Z"))
		})

		It("sets metadata", func() {
			subject.Metadata = []args.KeyValue{
				args.KeyValue{
					Key:   "key1",
					Value: "value1",
				},
				args.KeyValue{
					Key:   "key2",
					Value: "value2",
				},
			}

			err := subject.Execute([]string{})

			Expect(err).ToNot(HaveOccurred())
			Expect(storage.PutCallCount()).To(Equal(1))

			put := storage.PutArgsForCall(0)
			Expect(put.Metadata).To(HaveLen(2))
			Expect(put.Metadata[0].Key).To(Equal("key1"))
			Expect(put.Metadata[0].Value).To(Equal("value1"))
			Expect(put.Metadata[1].Key).To(Equal("key2"))
			Expect(put.Metadata[1].Value).To(Equal("value2"))
		})

		It("sets origin", func() {
			subject.Origin = []args.Origin{
				args.Origin{
					BlobReceiptOrigin: blobreceipt.BlobReceiptOrigin{
						"uri": "https://s3.amazonaws.com/bucket/key",
					},
				},
				args.Origin{
					BlobReceiptOrigin: blobreceipt.BlobReceiptOrigin{
						"uri":   "https://anotherbucket",
						"other": "fields",
					},
				},
			}

			err := subject.Execute([]string{})

			Expect(err).ToNot(HaveOccurred())
			Expect(storage.PutCallCount()).To(Equal(1))

			put := storage.PutArgsForCall(0)
			Expect(put.Origin).To(HaveLen(2))
			Expect(put.Origin[0]["uri"]).To(Equal("https://s3.amazonaws.com/bucket/key"))
			Expect(put.Origin[1]["uri"]).To(Equal("https://anotherbucket"))
			Expect(put.Origin[1]["other"]).To(Equal("fields"))
		})

		It("handles receipt write errors", func() {
			storage.PutReturns(errors.New("fake-err"))

			err := subject.Execute([]string{})

			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("Putting receipt"))
			Expect(err.Error()).To(ContainSubstring("fake-err"))
		})
	})
})
