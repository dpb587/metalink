package origin_test

import (
	. "github.com/dpb587/metalink/origin"

	boshsysfakes "github.com/cloudfoundry/bosh-utils/system/systemfakes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("DefaultFactory", func() {
	var subject OriginFactory
	var fs *boshsysfakes.FakeFileSystem

	BeforeEach(func() {
		fs = boshsysfakes.NewFakeFileSystem()

		subject = NewDefaultFactory(fs)
	})

	Describe("New", func() {
		It("handles files", func() {
			origin, err := subject.New("file:///some/tarmac")

			Expect(err).ToNot(HaveOccurred())
			Expect(origin.ReaderURI()).To(Equal("file:///some/tarmac"))
		})
	})
})
