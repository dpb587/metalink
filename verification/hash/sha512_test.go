package hash_test

import (
	"io/ioutil"
	"strings"

	"github.com/dpb587/metalink"
	"github.com/dpb587/metalink/file/filefakes"
	. "github.com/dpb587/metalink/verification/hash"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("SHA512Verification", func() {
	Describe("Signs", func() {
		It("hashes", func() {
			file := filefakes.FakeReference{}
			file.ReaderReturns(ioutil.NopCloser(strings.NewReader("the hash of crypt")), nil)

			result, err := SHA512SignerVerifier.Sign(&file)
			Expect(err).ToNot(HaveOccurred())

			meta4 := metalink.File{}

			err = result.Apply(&meta4)
			Expect(err).ToNot(HaveOccurred())

			Expect(meta4.Hashes).To(HaveLen(1))
			Expect(meta4.Hashes[0]).To(Equal(metalink.Hash{
				Type: "sha-512",
				Hash: "58a12ce8665e842168486fa7269d990e160e1f100f0dea9f7cb4b99789bc695b8923e4cc0663065868dfb7ade0d657362745101de76d9b77818375852e71eb22",
			}))
		})
	})
})
