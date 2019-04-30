package repositorypath_test

import (
	"github.com/dpb587/metalink/repository"
	. "github.com/dpb587/metalink/repository/filter/repositorypath"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Filter", func() {
	Describe("CreateFilter", func() {
		It("propagates", func() {
			filter, err := CreateFilter("something/prefix-*.meta4")
			Expect(err).NotTo(HaveOccurred())
			Expect(filter.Glob).To(Equal("something/prefix-*.meta4"))
		})
	})

	Describe("Filter", func() {
		Context("file wildcards", func() {
			It("matches", func() {
				match, err := Filter{Glob: "something/prefix-*.meta4"}.IsTrue(
					repository.RepositoryMetalink{
						Reference: repository.RepositoryMetalinkReference{
							Path: "something/prefix-v1.0.0.meta4",
						},
					},
				)

				Expect(err).NotTo(HaveOccurred())
				Expect(match).To(BeTrue())
			})

			It("does not match", func() {
				match, err := Filter{Glob: "something/prefix-*.meta4"}.IsTrue(
					repository.RepositoryMetalink{
						Reference: repository.RepositoryMetalinkReference{
							Path: "something/altfix-v1.0.0.meta4",
						},
					},
				)

				Expect(err).NotTo(HaveOccurred())
				Expect(match).To(BeFalse())
			})
		})

		Context("directory wildcards", func() {
			It("matches", func() {
				match, err := Filter{Glob: "prefix-*/else.meta4"}.IsTrue(
					repository.RepositoryMetalink{
						Reference: repository.RepositoryMetalinkReference{
							Path: "prefix-something/else.meta4",
						},
					},
				)

				Expect(err).NotTo(HaveOccurred())
				Expect(match).To(BeTrue())
			})

			It("does not match", func() {
				match, err := Filter{Glob: "prefix-*/else.meta4"}.IsTrue(
					repository.RepositoryMetalink{
						Reference: repository.RepositoryMetalinkReference{
							Path: "altfix-something/else.meta4",
						},
					},
				)

				Expect(err).NotTo(HaveOccurred())
				Expect(match).To(BeFalse())
			})
		})
	})
})
