package utility_test

import (
	"github.com/dpb587/metalink/repository/utility"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ParseUriOrGitCloneArg", func() {
	It("parses normal URLs", func() {
		uri, err := utility.ParseUriOrGitCloneArg("https://some.url/some-path")
		Expect(err).ToNot(HaveOccurred())
		Expect(uri.Scheme).To(Equal("https"))
		Expect(uri.Host).To(Equal("some.url"))
		Expect(uri.Path).To(Equal("/some-path"))
	})

	It("parses git+ssh clone-format strings", func() {
		uri, err := utility.ParseUriOrGitCloneArg("git+ssh://git@github.com:some-org/some-repo")
		Expect(err).ToNot(HaveOccurred())
		Expect(uri.Scheme).To(Equal("git+ssh"))
		Expect(uri.User.String()).To(Equal("git"))
		Expect(uri.Host).To(Equal("github.com"))
		Expect(uri.Path).To(Equal("some-org/some-repo"))
	})

	It("parses git+ssh clone-format strings without user", func() {
		uri, err := utility.ParseUriOrGitCloneArg("git+ssh://github.com:some-org/some-repo")
		Expect(err).ToNot(HaveOccurred())
		Expect(uri.Scheme).To(Equal("git+ssh"))
		Expect(uri.User.String()).To(Equal(""))
		Expect(uri.Host).To(Equal("github.com"))
		Expect(uri.Path).To(Equal("some-org/some-repo"))
	})

	It("returns an error for invalid URIs", func() {
		_, err := utility.ParseUriOrGitCloneArg("http://host:not-a-port")
		Expect(err).To(HaveOccurred())
	})
})
