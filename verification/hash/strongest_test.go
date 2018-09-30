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

var _ = Describe("StrongestSignerVerifier", func() {
	var file filefakes.FakeReference

	BeforeEach(func() {
		file = filefakes.FakeReference{}
		file.ReaderReturns(ioutil.NopCloser(strings.NewReader("the hash of crypt")), nil)
	})

	Describe("Sign", func() {
		It("prefers sha512", func() {
			verification, err := StrongestSignerVerifier.Sign(&file)
			Expect(err).ToNot(HaveOccurred())

			meta4 := metalink.File{}

			err = verification.Apply(&meta4)
			Expect(err).ToNot(HaveOccurred())

			Expect(meta4.Hashes).To(HaveLen(1))
			Expect(meta4.Hashes[0].Type).To(Equal(metalink.HashTypeSHA512))
		})
	})

	Describe("Verify", func() {
		It("errors when no hashes can be verified", func() {
			meta4 := metalink.File{
				Hashes: []metalink.Hash{
					{Type: "unknown", Hash: "bad"},
				},
			}

			result := StrongestSignerVerifier.Verify(&file, meta4)
			Expect(result.Error()).To(HaveOccurred())
		})

		It("prefers sha512", func() {
			meta4 := metalink.File{
				Hashes: []metalink.Hash{
					{Type: "sha-512", Hash: "58a12ce8665e842168486fa7269d990e160e1f100f0dea9f7cb4b99789bc695b8923e4cc0663065868dfb7ade0d657362745101de76d9b77818375852e71eb22"},
					{Type: "sha-256", Hash: "bad"},
					{Type: "sha-1", Hash: "bad"},
					{Type: "md5", Hash: "bad"},
				},
			}

			result := StrongestSignerVerifier.Verify(&file, meta4)
			Expect(result.Error()).ToNot(HaveOccurred())
			Expect(result.Verifier()).To(Equal("sha-512"))
			Expect(result.Confirmation()).To(Equal("OK"))
		})

		It("prefers sha256 without sha512", func() {
			meta4 := metalink.File{
				Hashes: []metalink.Hash{
					{Type: "sha-256", Hash: "bf7fee80eb7ee353b4af50fbe6decad4d73f3625d645e84cfd137935b50ea8dc"},
					{Type: "sha-1", Hash: "bad"},
					{Type: "md5", Hash: "bad"},
				},
			}

			result := StrongestSignerVerifier.Verify(&file, meta4)
			Expect(result.Error()).ToNot(HaveOccurred())
			Expect(result.Verifier()).To(Equal("sha-256"))
			Expect(result.Confirmation()).To(Equal("OK"))
		})

		It("prefers sha1 without sha256", func() {
			meta4 := metalink.File{
				Hashes: []metalink.Hash{
					{Type: "sha-1", Hash: "782e1a038874ebcd8877c7feb198d388e1b20569"},
					{Type: "md5", Hash: "bad"},
				},
			}

			result := StrongestSignerVerifier.Verify(&file, meta4)
			Expect(result.Error()).ToNot(HaveOccurred())
			Expect(result.Verifier()).To(Equal("sha-1"))
			Expect(result.Confirmation()).To(Equal("OK"))
		})

		It("prefers md5 without others", func() {
			meta4 := metalink.File{
				Hashes: []metalink.Hash{
					{Type: "md5", Hash: "05e24408c2048b675f812fbffcfba4ea"},
				},
			}

			result := StrongestSignerVerifier.Verify(&file, meta4)
			Expect(result.Error()).ToNot(HaveOccurred())
			Expect(result.Verifier()).To(Equal("md5"))
			Expect(result.Confirmation()).To(Equal("OK"))
		})
	})
})
