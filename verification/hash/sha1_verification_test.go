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

var _ = Describe("SHA1Verification", func() {
	Describe("Signs", func() {
		It("hashes", func() {
			file := filefakes.FakeReference{}
			file.ReaderReturns(ioutil.NopCloser(strings.NewReader("the hash of crypt")), nil)

			result, err := SHA1Verification.Sign(&file)
			Expect(err).ToNot(HaveOccurred())

			meta4 := metalink.File{}

			err = result.Apply(&meta4)
			Expect(err).ToNot(HaveOccurred())

			Expect(meta4.Hashes).To(HaveLen(1))
			Expect(meta4.Hashes[0]).To(Equal(metalink.Hash{
				Type: "sha-1",
				Hash: "782e1a038874ebcd8877c7feb198d388e1b20569",
			}))
		})
	})
})
