package file_test

import (
	"io/ioutil"
	"os"
	"path"

	"github.com/dpb587/metalink/file"
	. "github.com/dpb587/metalink/file/url/file"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("File", func() {
	var subject file.Reference
	var tmpfile *os.File

	BeforeEach(func() {
		var err error

		tmpfile, err = ioutil.TempFile("", "boshua-file-test-")
		Expect(err).NotTo(HaveOccurred())
	})

	AfterEach(func() {
		Expect(os.Remove(tmpfile.Name())).To(Succeed())
	})

	Describe("NewReference", func() {
		XIt("expands paths", func() {
			subject := NewReference("~/somewhere")

			Expect(subject).ToNot(BeNil())
			Expect(subject.ReaderURI()).To(Equal("file:///root/somewhere"))
		})
	})

	Describe("Name", func() {
		It("gives the base name", func() {
			subject = NewReference(tmpfile.Name())

			value, err := subject.Name()

			Expect(err).ToNot(HaveOccurred())
			Expect(value).To(Equal(path.Base(tmpfile.Name())))
		})
	})

	Describe("Size", func() {
		It("gives the size", func() {
			tmpfile.Write([]byte("something useful"))

			subject = NewReference(tmpfile.Name())

			value, err := subject.Size()

			Expect(err).ToNot(HaveOccurred())
			Expect(value).To(Equal(uint64(16)))
		})

		It("errors gracefully", func() {
			subject = NewReference("/fake/path")

			_, err := subject.Size()

			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("Checking file size"))
		})
	})

	Describe("Reader", func() {
		It("opens for reading", func() {
			tmpfile.Write([]byte("something useful"))

			subject = NewReference(tmpfile.Name())

			reader, err := subject.Reader()
			Expect(err).ToNot(HaveOccurred())

			readerString, _ := ioutil.ReadAll(reader)
			Expect(string(readerString)).To(Equal("something useful"))
		})

		It("errors gracefully", func() {
			subject = NewReference("/fake/path")

			_, err := subject.Reader()

			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("Opening file"))
		})
	})
})
