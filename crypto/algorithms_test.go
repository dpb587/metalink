package crypto_test

import (
	. "github.com/dpb587/blob-receipt/crypto"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	boshcry "github.com/cloudfoundry/bosh-utils/crypto"
)

var _ = Describe("Algorithms", func() {
	Describe("GetAlgorithms", func() {
		It("knows md5", func() {
			algorithm, err := GetAlgorithm("md5")

			Expect(err).ToNot(HaveOccurred())
			Expect(algorithm.Name()).To(Equal("md5"))
		})

		It("knows sha1", func() {
			algorithm, err := GetAlgorithm("sha1")

			Expect(err).ToNot(HaveOccurred())
			Expect(algorithm).ToNot(BeNil())
			Expect(algorithm.Name()).To(Equal("sha1"))
		})

		It("knows sha256", func() {
			algorithm, err := GetAlgorithm("sha256")

			Expect(err).ToNot(HaveOccurred())
			Expect(algorithm).ToNot(BeNil())
			Expect(algorithm.Name()).To(Equal("sha256"))
		})

		It("knows sha512", func() {
			algorithm, err := GetAlgorithm("sha512")

			Expect(err).ToNot(HaveOccurred())
			Expect(algorithm).ToNot(BeNil())
			Expect(algorithm.Name()).To(Equal("sha512"))
		})

		It("errors on unknown", func() {
			_, err := GetAlgorithm("superhash")

			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("Unsupported digest algorithm: superhash"))
		})
	})

	Describe("GetStrongestAlgorithm", func() {
		It("prefers sha512 over sha256", func() {
			algorithm, err := GetStrongestAlgorithm([]string{"sha256", "sha512"})

			Expect(err).ToNot(HaveOccurred())
			Expect(algorithm).ToNot(BeNil())
			Expect(algorithm.Name()).To(Equal("sha512"))
		})

		It("prefers sha256 over sha1", func() {
			algorithm, err := GetStrongestAlgorithm([]string{"sha1", "sha256"})

			Expect(err).ToNot(HaveOccurred())
			Expect(algorithm).ToNot(BeNil())
			Expect(algorithm.Name()).To(Equal("sha256"))
		})

		It("prefers sha1 over md5", func() {
			algorithm, err := GetStrongestAlgorithm([]string{"md5", "sha1"})

			Expect(err).ToNot(HaveOccurred())
			Expect(algorithm).ToNot(BeNil())
			Expect(algorithm.Name()).To(Equal("sha1"))
		})

		It("prefers md5 over unknown", func() {
			algorithm, err := GetStrongestAlgorithm([]string{"superhash", "md5"})

			Expect(err).ToNot(HaveOccurred())
			Expect(algorithm).ToNot(BeNil())
			Expect(algorithm.Name()).To(Equal("md5"))
		})

		It("errors with unknown algorithms", func() {
			_, err := GetStrongestAlgorithm([]string{"superhash", "unsuperhash"})

			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("No strong algorithm found: superhash, unsuperhash"))
		})
	})

	Describe("GetDigestHash", func() {
		It("strips prefixes", func() {
			digestHash := GetDigestHash(boshcry.NewDigest(boshcry.DigestAlgorithmSHA512, "fake-hash"))

			Expect(digestHash).To(Equal("fake-hash"))
		})

		It("ignores sha1 non-prefixes", func() {
			digestHash := GetDigestHash(boshcry.NewDigest(boshcry.DigestAlgorithmSHA1, "sha1:fake-hash"))

			Expect(digestHash).To(Equal("sha1:fake-hash"))
		})
	})
})
