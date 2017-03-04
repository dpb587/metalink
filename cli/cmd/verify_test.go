package cmd_test

import (
	"errors"

	blobreceipt "github.com/dpb587/blob-receipt"
	. "github.com/dpb587/blob-receipt/cli/cmd"
	"github.com/dpb587/blob-receipt/origin/originfakes"
	"github.com/dpb587/blob-receipt/storage/storagefakes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Verify", func() {
	var subject Verify
	var originFactory originfakes.FakeOriginFactory
	var storageFactory storagefakes.FakeStorageFactory

	BeforeEach(func() {
		originFactory = originfakes.FakeOriginFactory{}
		storageFactory = storagefakes.FakeStorageFactory{}

		subject = Verify{
			Args:           VerifyArgs{},
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

		It("handles storage errors", func() {
			storage.GetReturns(blobreceipt.BlobReceipt{}, errors.New("fake-err"))

			err := subject.Execute([]string{})

			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("Loading receipt"))
			Expect(err.Error()).To(ContainSubstring("fake-err"))
		})
	})
})
