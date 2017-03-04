package origin_test

import (
	"errors"
	"io/ioutil"
	"time"

	. "github.com/dpb587/blob-receipt/origin"

	boshcry "github.com/cloudfoundry/bosh-utils/crypto"
	boshsysfakes "github.com/cloudfoundry/bosh-utils/system/systemfakes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("FileSystem", func() {
	var subject Origin
	var fs *boshsysfakes.FakeFileSystem

	BeforeEach(func() {
		fs = boshsysfakes.NewFakeFileSystem()

		fs.WriteFileString("/somewhere/useful", "something useful")
	})

	Describe("CreateFileSystem", func() {
		It("expands paths", func() {
			fs.ExpandPathExpanded = "/root/somewhere"

			subject, err := CreateFileSystem(fs, "~/somewhere")

			Expect(err).ToNot(HaveOccurred())
			Expect(subject).ToNot(BeNil())
			Expect(subject.String()).To(Equal("/root/somewhere"))
		})

		It("errors when expansion fails", func() {
			fs.ExpandPathErr = errors.New("fake-err-1")

			_, err := CreateFileSystem(fs, "~/somewhere")

			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("Expanding path"))
		})
	})

	Describe("Digest", func() {
		It("opens for reading", func() {
			subject, _ = CreateFileSystem(fs, "/somewhere/useful")

			digest, err := subject.Digest(boshcry.DigestAlgorithmSHA1)

			Expect(err).ToNot(HaveOccurred())
			Expect(digest).ToNot(BeNil())
			Expect(digest.String()).To(Equal("6be5595fd5000d677e6543bd5bbead66371951ec"))
		})

		It("errors gracefully", func() {
			fs.OpenFileErr = errors.New("fake-err")

			subject, _ = CreateFileSystem(fs, "/somewhere/useful")

			_, err := subject.Digest(boshcry.DigestAlgorithmSHA1)

			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("Digesting file"))
			Expect(err.Error()).To(ContainSubstring("fake-err"))
		})
	})

	Describe("Name", func() {
		It("gives the base name", func() {
			subject, _ = CreateFileSystem(fs, "/somewhere/useful")

			value, err := subject.Name()

			Expect(err).ToNot(HaveOccurred())
			Expect(value).To(Equal("useful"))
		})
	})

	Describe("Size", func() {
		It("gives the size", func() {
			subject, _ = CreateFileSystem(fs, "/somewhere/useful")

			value, err := subject.Size()

			Expect(err).ToNot(HaveOccurred())
			Expect(value).To(Equal(uint64(16)))
		})

		It("errors gracefully", func() {
			fs.RegisterOpenFile("/somewhere/useful", &boshsysfakes.FakeFile{
				StatErr: errors.New("fake-err"),
			})

			subject, _ = CreateFileSystem(fs, "/somewhere/useful")

			_, err := subject.Size()

			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("Checking file size"))
		})
	})

	XDescribe("Time", func() {
		// @todo fakefs doesn't seem to track time
		It("gives the time", func() {
			subject, _ = CreateFileSystem(fs, "/somewhere/useful")

			value, err := subject.Time()

			Expect(err).ToNot(HaveOccurred())
			// @todo improve?
			Expect(value).To(BeNumerically(">", time.Now().Add(-5*time.Second)))
			Expect(value).To(BeNumerically("<", time.Now().Add(5*time.Second)))
		})

		It("errors gracefully", func() {
			fs.RegisterOpenFile("/somewhere/useful", &boshsysfakes.FakeFile{
				StatErr: errors.New("fake-err"),
			})

			subject, _ = CreateFileSystem(fs, "/somewhere/useful")

			_, err := subject.Time()

			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("Checking file time"))
		})
	})

	Describe("Reader", func() {
		It("opens for reading", func() {
			subject, _ = CreateFileSystem(fs, "/somewhere/useful")

			reader, err := subject.Reader()

			Expect(err).ToNot(HaveOccurred())

			readerString, _ := ioutil.ReadAll(reader)

			Expect(string(readerString)).To(Equal("something useful"))
		})

		It("errors gracefully", func() {
			fs.OpenFileErr = errors.New("fake-err")

			subject, _ = CreateFileSystem(fs, "/somewhere/useful")

			_, err := subject.Reader()

			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("Opening file"))
			Expect(err.Error()).To(ContainSubstring("fake-err"))
		})
	})
})
