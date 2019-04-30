package template_test

import (
	"github.com/dpb587/metalink"

	. "github.com/dpb587/metalink/template"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("Template", func() {
	var file metalink.File

	BeforeEach(func() {
		file = metalink.File{
			Name: "fake-name",
			Hashes: []metalink.Hash{
				{
					Type: "md5",
					Hash: "fake-md5",
				},
				{
					Type: "sha-1",
					Hash: "fake-sha-1",
				},
				{
					Type: "sha-256",
					Hash: "fake-sha-256",
				},
				{
					Type: "sha-512",
					Hash: "fake-sha-512",
				},
			},
		}
	})

	DescribeTable(
		"hash value when present",
		func(uri string, expected string) {
			template, err := New(uri)
			Expect(err).NotTo(HaveOccurred())

			res, err := template.ExecuteString(file)
			Expect(err).NotTo(HaveOccurred())
			Expect(res).To(Equal(expected))
		},
		Entry("md5", "{{ .MD5 }}", "fake-md5"),
		Entry("sha1", "{{ .SHA1 }}", "fake-sha-1"),
		Entry("sha256", "{{ .SHA256 }}", "fake-sha-256"),
		Entry("sha512", "{{ .SHA512 }}", "fake-sha-512"),
	)

	DescribeTable(
		"empty when missing",
		func(uri string) {
			template, err := New(uri)
			Expect(err).NotTo(HaveOccurred())

			res, err := template.ExecuteString(metalink.File{})
			Expect(err).NotTo(HaveOccurred())
			Expect(res).To(Equal(""))
		},
		Entry("md5", "{{ .MD5 }}"),
		Entry("sha1", "{{ .SHA1 }}"),
		Entry("sha256", "{{ .SHA256 }}"),
		Entry("sha512", "{{ .SHA512 }}"),
	)

	It("supports regular fields", func() {
		template, err := New("{{ .Name }}")
		Expect(err).NotTo(HaveOccurred())

		res, err := template.ExecuteString(metalink.File{Name: "fake-name"})
		Expect(err).NotTo(HaveOccurred())
		Expect(res).To(Equal("fake-name"))
	})

	It("errors for bad template parses", func() {
		_, err := New("{{ oops")
		Expect(err).To(HaveOccurred())
	})

	It("errors for bad template parses", func() {
		template, err := New("{{ .MissingField }}")
		Expect(err).NotTo(HaveOccurred())

		_, err = template.ExecuteString(metalink.File{})
		Expect(err).To(HaveOccurred())
	})
})
