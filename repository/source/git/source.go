package git

import (
	"crypto/md5"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"strings"

	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	boshsys "github.com/cloudfoundry/bosh-utils/system"
	"github.com/dpb587/metalink"
	"github.com/dpb587/metalink/repository"
	"github.com/dpb587/metalink/repository/filter"
	"github.com/dpb587/metalink/repository/source"
)

type Source struct {
	rawURI     string
	uri        string
	branch     string
	path       string
	privateKey *string
	commits    sourceCommitSettings
	fs         boshsys.FileSystem
	cmdRunner  boshsys.CmdRunner

	clonegit string
	clonedir string

	metalinks []repository.RepositoryMetalink
}

type sourceCommitSettings struct {
	committerName  string
	committerEmail string
	authorName     string
	authorEmail    string
	message        string
}

var _ source.Source = &Source{}

func NewSource(rawURI string, uri string, branch string, path string, privateKey *string, commits sourceCommitSettings, fs boshsys.FileSystem, cmdRunner boshsys.CmdRunner) *Source {
	return &Source{
		rawURI:     rawURI,
		uri:        uri,
		branch:     branch,
		path:       path,
		privateKey: privateKey,
		commits:    commits,
		fs:         fs,
		cmdRunner:  cmdRunner,
	}
}

func (s *Source) Load() error {
	err := s.requireClone()
	if err != nil {
		return bosherr.WrapError(err, "Cloning repository")
	}

	files, err := s.fs.Glob(fmt.Sprintf("%s/%s/*.meta4", s.clonedir, s.path))
	if err != nil {
		return bosherr.WrapError(err, "Listing metalinks")
	}

	uri := s.URI()
	s.metalinks = []repository.RepositoryMetalink{}

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
			WorkingDir: s.clonedir,
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

		repometa4 := repository.RepositoryMetalink{
			Reference: repository.RepositoryMetalinkReference{
				Repository: uri,
				Path:       strings.TrimPrefix(file, fmt.Sprintf("%s/%s/", s.clonedir, s.path)),
				Version:    version,
			},
		}

		err = metalink.Unmarshal(metalinkBytes, &repometa4.Metalink)
		if err != nil {
			return bosherr.WrapError(err, "Unmarshaling")
		}

		s.metalinks = append(s.metalinks, repometa4)
	}

	return nil
}

func (s Source) URI() string {
	return s.rawURI
}

func (s Source) Filter(f filter.Filter) ([]repository.RepositoryMetalink, error) {
	return source.FilterInMemory(s.metalinks, f)
}

func (s Source) Put(name string, data io.Reader) error {
	err := s.requireClone()
	if err != nil {
		return bosherr.WrapError(err, "Cloning repository")
	}

	filepath := path.Join(s.path, name)

	content, err := ioutil.ReadAll(data)
	if err != nil {
		return bosherr.WrapError(err, "Reading metalink")
	}

	err = s.fs.WriteFile(path.Join(s.clonedir, filepath), content)
	if err != nil {
		return bosherr.WrapError(err, "Writing metalink")
	}

	_, _, exitStatus, err := s.cmdRunner.RunComplexCommand(boshsys.Command{
		Name:       "git",
		Args:       []string{"add", filepath},
		WorkingDir: s.clonedir,
	})
	if err != nil {
		return bosherr.WrapError(err, "Staging metalink")
	} else if exitStatus != 0 {
		return fmt.Errorf("git add exit status: %d", exitStatus)
	}

	_, _, exitStatus, err = s.cmdRunner.RunComplexCommand(boshsys.Command{
		Name: "git",
		Args: []string{
			"commit",
			"-m",
			s.commits.message,
			filepath,
		},
		WorkingDir: s.clonedir,
		Env: map[string]string{
			"GIT_AUTHOR_EMAIL":    s.commits.authorEmail,
			"GIT_AUTHOR_NAME":     s.commits.authorName,
			"GIT_COMMITTER_EMAIL": s.commits.committerEmail,
			"GIT_COMMITTER_NAME":  s.commits.committerName,
		},
	})
	if err != nil {
		return bosherr.WrapError(err, "Creating commit")
	} else if exitStatus != 0 {
		return fmt.Errorf("git commit exit status: %d", exitStatus)
	}

	_, _, exitStatus, err = s.cmdRunner.RunComplexCommand(boshsys.Command{
		Name:       s.clonegit,
		Args:       []string{"push"},
		WorkingDir: s.clonedir,
	})
	if err != nil {
		return bosherr.WrapError(err, "Pushing repository")
	} else if exitStatus != 0 {
		return fmt.Errorf("git push exit status: %d", exitStatus)
	}

	return nil
}

func (s *Source) requireClone() error {
	if s.clonedir != "" {
		return nil
	}

	tmpdir := fmt.Sprintf("%s/metalink-git-source-%x-1", strings.TrimSuffix(os.TempDir(), "/"), md5.Sum([]byte(s.rawURI)))

	err := s.fs.MkdirAll(tmpdir, 0700)
	if err != nil {
		return bosherr.WrapError(err, "Creating tmpdir for git")
	}

	if s.privateKey == nil {
		s.clonegit = "git"
	} else {
		s.clonegit = fmt.Sprintf("%s.git", tmpdir)

		keyPath := fmt.Sprintf("%s.key", s.clonegit)

		err = s.fs.WriteFile(keyPath, []byte(*s.privateKey))
		if err != nil {
			return bosherr.WrapError(err, "Writing private key")
		}

		err = s.fs.Chmod(keyPath, 0600)
		if err != nil {
			return bosherr.WrapError(err, "Securing private key")
		}

		err = s.fs.WriteFileString(s.clonegit, fmt.Sprintf(`#!/bin/bash
eval $(ssh-agent)
trap "kill $SSH_AGENT_PID" 0
set -eu
SSH_ASKPASS=false DISPLAY= ssh-add "%s"
git "$@"
`, keyPath))
		if err != nil {
			return bosherr.WrapError(err, "Writing git wrapper")
		}

		err = s.fs.Chmod(s.clonegit, 0755)
		if err != nil {
			return bosherr.WrapError(err, "Chmod'ing git wrapper")
		}
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
			Name:       s.clonegit,
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

		_, _, exitStatus, err := s.cmdRunner.RunCommand(s.clonegit, args...)
		if err != nil {
			return bosherr.WrapError(err, "Cloning repository")
		} else if exitStatus != 0 {
			return fmt.Errorf("git clone exit status: %d", exitStatus)
		}
	}

	s.clonedir = tmpdir

	return nil
}
