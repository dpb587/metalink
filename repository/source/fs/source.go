package fs

import (
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"

	boshsys "github.com/cloudfoundry/bosh-utils/system"
	"github.com/dpb587/metalink"
	"github.com/dpb587/metalink/repository"
	"github.com/dpb587/metalink/repository/filter"
	"github.com/dpb587/metalink/repository/source"
	"github.com/pkg/errors"
)

type Source struct {
	uri  string
	fs   boshsys.FileSystem
	path string

	metalinks []repository.RepositoryMetalink
}

var _ source.Source = &Source{}

func NewSource(uri string, fs boshsys.FileSystem, path string) *Source {
	return &Source{
		uri:  uri,
		fs:   fs,
		path: path,
	}
}

func (s *Source) Load() error {
	uri := s.URI()
	s.metalinks = []repository.RepositoryMetalink{}

	legacyPaths, err := filepath.Glob(s.path)
	if err != nil {
		return errors.Wrap(err, "Globbing path")
	}

	for _, legacyPath := range legacyPaths {
		var files []string

		err := s.fs.Walk(legacyPath, func(p string, _ os.FileInfo, err error) error {
			if err != nil {
				return err
			} else if path.Ext(p) != ".meta4" {
				return nil
			}

			files = append(files, p)

			return nil
		})
		if err != nil {
			return errors.Wrap(err, "Walking path")
		}

		for _, file := range files {
			metalinkBytes, err := s.fs.ReadFile(file)
			if err != nil {
				return errors.Wrap(err, "Reading metalink")
			}

			repometa4 := repository.RepositoryMetalink{
				Reference: repository.RepositoryMetalinkReference{
					Repository: uri,
					Path:       strings.TrimPrefix(strings.TrimPrefix(file, legacyPath), "/"),
				},
			}

			err = metalink.Unmarshal(metalinkBytes, &repometa4.Metalink)
			if err != nil {
				return errors.Wrap(err, "Unmarshaling")
			}

			s.metalinks = append(s.metalinks, repometa4)
		}
	}

	return nil
}

func (s Source) URI() string {
	return s.uri
}

func (s Source) Filter(f filter.Filter) ([]repository.RepositoryMetalink, error) {
	return source.FilterInMemory(s.metalinks, f)
}

func (s Source) Put(name string, data io.Reader) error {
	path := path.Join(s.path, name)

	content, err := ioutil.ReadAll(data)
	if err != nil {
		return errors.Wrap(err, "Reading metalink")
	}

	err = s.fs.WriteFile(path, content)
	if err != nil {
		return errors.Wrap(err, "Writing metalink")
	}

	return nil
}
