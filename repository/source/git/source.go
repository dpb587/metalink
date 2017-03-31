package git

import (
	"crypto/md5"
	"fmt"
	"os"
	"strings"

	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	boshsys "github.com/cloudfoundry/bosh-utils/system"
	"github.com/dpb587/metalink/repository"
	"github.com/dpb587/metalink/repository/filter"
	"github.com/dpb587/metalink/repository/source"
)

type Source struct {
	rawURI    string
	uri       string
	branch    string
	path      string
	fs        boshsys.FileSystem
	cmdRunner boshsys.CmdRunner

	files []repository.File
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

func (s *Source) Load() error {
	tmpdir := fmt.Sprintf("%s/metalink-git-source-%x-1", os.TempDir(), md5.Sum([]byte(s.rawURI)))

	err := s.fs.MkdirAll(tmpdir, 0700)
	if err != nil {
		return bosherr.WrapError(err, "Creating tmpdir for git")
	}

	if s.fs.FileExists(fmt.Sprintf("%s/.git", tmpdir)) {
		args := []string{
			"pull",
			"--ff-only",
			s.uri,
		}

		if s.branch != "" {
			args = append(args, "--branch", s.branch)
		}

		_, _, exitStatus, err := s.cmdRunner.RunComplexCommand(boshsys.Command{
			Name:       "git",
			Args:       args,
			WorkingDir: tmpdir,
		})
		if err != nil {
			return bosherr.WrapError(err, "Pulling repository")
		} else if exitStatus != 0 {
			return fmt.Errorf("git pull exit status: %d", exitStatus)
		}
	} else {
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
	}

	files, err := s.fs.Glob(fmt.Sprintf("%s/%s/*.meta4", tmpdir, s.path))
	if err != nil {
		return bosherr.WrapError(err, "Listing metalinks")
	}

	s.files = []repository.File{}

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

		metalinkBytes, err := s.fs.ReadFile(file)
		if err != nil {
			return bosherr.WrapError(err, "Reading metalink")
		}

		results, err := source.ExplodeMetalinkBytes(
			repository.Repository{
				URI:     s.URI(),
				Path:    strings.TrimPrefix(file, fmt.Sprintf("%s/%s/", tmpdir, s.path)),
				Version: version,
			},
			metalinkBytes,
		)
		if err != nil {
			return bosherr.WrapError(err, "Loading metalink")
		}

		s.files = append(s.files, results...)
	}

	return nil
}

func (s Source) URI() string {
	return s.rawURI
}

func (s Source) FilterFiles(filter filter.Filter) ([]repository.File, error) {
	return source.FilterFilesInMemory(s.files, filter)
}
