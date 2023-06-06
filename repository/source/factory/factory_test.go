package factory_test

import (
	boshlog "github.com/cloudfoundry/bosh-utils/logger"
	boshsys "github.com/cloudfoundry/bosh-utils/system"
	"github.com/dpb587/metalink/repository/source/factory"
	source_fs "github.com/dpb587/metalink/repository/source/fs"
	source_git "github.com/dpb587/metalink/repository/source/git"
	source_http "github.com/dpb587/metalink/repository/source/http"
	source_s3 "github.com/dpb587/metalink/repository/source/s3"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Factory", func() {
	var subject *factory.Factory

	Describe("Create", func() {
		BeforeEach(func() {
			logger := boshlog.NewLogger(boshlog.LevelError)
			fs := boshsys.NewOsFileSystem(logger)
			cmdRunner := boshsys.NewExecCmdRunner(logger)

			subject = factory.NewFactory()
			subject.Add(source_fs.NewFactory(fs))
			subject.Add(source_http.NewFactory())
			subject.Add(source_git.NewFactory(fs, cmdRunner))
			subject.Add(source_s3.NewFactory())
		})

		It("supports file URLs", func() {
			source, err := subject.Create("file:///some/file/path", map[string]interface{}{})
			Expect(err).ToNot(HaveOccurred())
			_, isCorrectType := source.(*source_fs.Source)
			Expect(isCorrectType).To(BeTrue())
		})

		It("supports HTTP URLs", func() {
			source, err := subject.Create("https://some.url", map[string]interface{}{})
			Expect(err).ToNot(HaveOccurred())
			_, isCorrectType := source.(*source_http.Source)
			Expect(isCorrectType).To(BeTrue())
		})

		It("supports S3 URLs", func() {
			source, err := subject.Create("s3://s3.amazonaws.com/some-bucket/some-file", map[string]interface{}{})
			Expect(err).ToNot(HaveOccurred())
			_, isCorrectType := source.(*source_s3.Source)
			Expect(isCorrectType).To(BeTrue())
		})

		It("supports git SSH clone-format URLs", func() {
			source, err := subject.Create("git+ssh://git@github.com:some-org/some-repo", map[string]interface{}{})
			Expect(err).ToNot(HaveOccurred())
			_, isCorrectType := source.(*source_git.Source)
			Expect(isCorrectType).To(BeTrue())
		})

		It("supports git SSH URLs", func() {
			source, err := subject.Create("git+ssh://github.com/some-org/some-repo", map[string]interface{}{})
			Expect(err).ToNot(HaveOccurred())
			_, isCorrectType := source.(*source_git.Source)
			Expect(isCorrectType).To(BeTrue())
		})
	})
})
