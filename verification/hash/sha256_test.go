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

var _ = Describe("SHA256Verification", func() {
	Describe("Signs", func() {
		It("hashes", func() {
			file := filefakes.FakeReference{}
			file.ReaderReturns(ioutil.NopCloser(strings.NewReader("the hash of crypt")), nil)

			result, err := SHA256SignerVerifier.Sign(&file)
			Expect(err).ToNot(HaveOccurred())

			meta4 := metalink.File{}

			err = result.Apply(&meta4)
			Expect(err).ToNot(HaveOccurred())

			Expect(meta4.Hashes).To(HaveLen(1))
			Expect(meta4.Hashes[0]).To(Equal(metalink.Hash{
				Type: "sha-256",
				Hash: "bf7fee80eb7ee353b4af50fbe6decad4d73f3625d645e84cfd137935b50ea8dc",
			}))
		})
	})
})
