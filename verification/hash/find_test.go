package hash_test

import (
	"github.com/dpb587/metalink"
	. "github.com/dpb587/metalink/verification/hash"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Find", func() {
	It("finds hash when present", func() {
		hash, found := Find(
			metalink.File{
				Hashes: []metalink.Hash{
					{
						Type: "not-sha-1",
						Hash: "not-sha-1-hash",
					},
					{
						Type: "sha-1",
						Hash: "sha-1-hash",
					},
				},
			},
			"sha-1",
		)

		Expect(found).To(BeTrue())
		Expect(hash).To(Equal(metalink.Hash{
			Type: "sha-1",
			Hash: "sha-1-hash",
		}))
	})

	It("returns nothing when not found", func() {
		hash, found := Find(
			metalink.File{
				Hashes: []metalink.Hash{
					{
						Type: "sha-1",
						Hash: "sha-1-hash",
					},
				},
			},
			"sha-256",
		)

		Expect(found).To(BeFalse())
		Expect(hash).To(Equal(metalink.Hash{}))
	})
})
