package git

import (
	"encoding/json"
	"fmt"
	"path"

	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	boshsys "github.com/cloudfoundry/bosh-utils/system"
	"github.com/dpb587/blob-receipt"
	"github.com/dpb587/blob-receipt/repository"
	"github.com/dpb587/blob-receipt/repository/filter"
	"github.com/dpb587/blob-receipt/repository/source"
)

type Source struct {
	rawURI    string
	uri       string
	branch    string
	path      string
	fs        boshsys.FileSystem
	cmdRunner boshsys.CmdRunner

	receipts []repository.BlobReceipt
}

var _ source.Source = &Source{}

func NewSource(rawURI string, uri string, branch string, path string, fs boshsys.FileSystem, cmdRunner boshsys.CmdRunner) *Source {
	return &Source{
		rawURI:    rawURI,
		uri:       uri,
		branch:    branch,
		path:      path,
		fs:        fs,
		cmdRunner: cmdRunner,
	}
}

func (s *Source) Reload() error {
	tmpdir, err := s.fs.TempDir("blob-git")
	if err != nil {
		return bosherr.WrapError(err, "Creating tmpdir for git")
	}

	defer s.fs.RemoveAll(tmpdir)

	args := []string{
		"clone",
		"--single-branch",
	}

	if s.branch != "" {
		args = append(args, "--branch", s.branch)
	}

	args = append(args, s.uri, tmpdir)

	_, _, exitStatus, err := s.cmdRunner.RunCommand("git", args...)
	if err != nil {
		return bosherr.WrapError(err, "Cloning repository")
	} else if exitStatus != 0 {
		return fmt.Errorf("git clone exit status: %d", exitStatus)
	}

	files, err := s.fs.Glob(fmt.Sprintf("%s/%s/*.json", tmpdir, s.path))
	if err != nil {
		return bosherr.WrapError(err, "Listing receipts")
	}

	s.receipts = []repository.BlobReceipt{}

	for _, file := range files {
		command := boshsys.Command{
			Name: "git",
			Args: []string{
				"log",
				"--pretty=format:%H",
				"-n1",
				"--",
				file,
			},
			WorkingDir: tmpdir,
		}

		version, _, exitStatus, err := s.cmdRunner.RunComplexCommand(command)
		if err != nil {
			return bosherr.WrapError(err, "Getting version of file")
		} else if exitStatus != 0 {
			return fmt.Errorf("git log exit status: %d", exitStatus)
		}

		receiptBytes, err := s.fs.ReadFile(file)
		if err != nil {
			return bosherr.WrapError(err, "Reading receipt")
		}

		receipt := blobreceipt.BlobReceipt{}

		err = json.Unmarshal(receiptBytes, &receipt)
		if err != nil {
			return bosherr.WrapError(err, "Parsing receipt")
		}

		annotatedreceipt := repository.BlobReceipt{
			Repository: repository.BlobReceiptRepository{
				URI:     s.URI(),
				Path:    path.Base(file),
				Version: version,
			},
			Receipt: receipt,
		}

		s.receipts = append(s.receipts, annotatedreceipt)
	}

	return nil
}

func (s Source) URI() string {
	return s.rawURI
}

func (s Source) FilterBlobReceipts(filter filter.Filter) ([]repository.BlobReceipt, error) {
	matches := []repository.BlobReceipt{}

	for _, receipt := range s.receipts {
		matched, err := filter.IsTrue(receipt.Receipt)
		if err != nil {
			return nil, bosherr.WrapError(err, "Matching receipt")
		} else if !matched {
			continue
		}

		matches = append(matches, receipt)
	}

	return matches, nil
}
