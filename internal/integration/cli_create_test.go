package integration_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/onsi/gomega/gexec"
)

var _ = Describe("create", func() {
	var executable string
	var tmpdir string
	var blobfile string
	var receiptfile string

	BeforeSuite(func() {
		var err error

		executable, err = gexec.Build("github.com/dpb587/blob-receipt/cli")
		Expect(err).ToNot(HaveOccurred())
	})

	BeforeEach(func() {
		var err error

		tmpdir, err = ioutil.TempDir(os.TempDir(), "cli-create")
		Expect(err).ToNot(HaveOccurred())

		blobfile = fmt.Sprintf("%s/blob", tmpdir)
		receiptfile = fmt.Sprintf("%s/receipt", tmpdir)

		file, err := os.OpenFile(blobfile, os.O_CREATE|os.O_WRONLY, 0700)
		Expect(err).ToNot(HaveOccurred())
		defer file.Close()

		_, err = file.WriteString(`dummy path b6925`)
		Expect(err).ToNot(HaveOccurred())
	})

	AfterEach(func() {
		os.RemoveAll(tmpdir)
	})

	AfterSuite(func() {
		os.Remove(executable)
	})

	Describe("basics", func() {
		It("executes", func() {
			cmd := exec.Command(
				executable,
				"create",
				"--time", "2017-01-02T03:04:05Z",
				"--metadata", "key1=value1",
				"--metadata", "key2=value2=value2",
				"--origin", "https://blobs.example.com/source-angeles",
				"--origin", `{"uri":"https://blobs2.example.com/source-angeles","extra":"backup"}`,
				receiptfile,
				blobfile,
			)

			session, err := gexec.Start(cmd, GinkgoWriter, GinkgoWriter)
			Expect(err).ToNot(HaveOccurred())

			session.Wait(2 * time.Second)

			Expect(session.ExitCode()).To(Equal(0))

			file, err := os.Open(receiptfile)
			Expect(err).ToNot(HaveOccurred())

			receiptcontents, err := ioutil.ReadAll(file)
			Expect(err).ToNot(HaveOccurred())

			Expect(string(receiptcontents)).To(ContainSubstring(`{
  "digest": {
    "md5": "abc04197c7d725e8befce32154cf7dfa",
    "sha1": "3582a49c84025af672152273bd9818226b62b1f6",
    "sha256": "bfc83a441083e276df5ad8be431ea9de3481bce56226cbd4d77134eb3c797105",
    "sha512": "d1b7299881b842b319dd00133462aa570a57eb80506b3fd661aa70a16918653ba4912de5ddfe38b1116084de8c9e09b6d55c450c3b624e791506d617632a2a75"
  },
  "metadata": [
    {
      "key": "key1",
      "value": "value1"
    },
    {
      "key": "key2",
      "value": "value2=value2"
    }
  ],
  "name": "blob",
  "origin": [
    {
      "uri": "https://blobs.example.com/source-angeles"
    },
    {
      "extra": "backup",
      "uri": "https://blobs2.example.com/source-angeles"
    }
  ],
  "size": 16,
  "time": "2017-01-02T03:04:05Z"
}
`))
		})
	})
})
