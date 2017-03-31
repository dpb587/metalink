package fs

import (
	"fmt"
	"path"
	"time"

	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	boshsys "github.com/cloudfoundry/bosh-utils/system"
	"github.com/dpb587/metalink/repository"
	"github.com/dpb587/metalink/repository/filter"
	"github.com/dpb587/metalink/repository/source"
)

type Source struct {
	uri  string
	fs   boshsys.FileSystem
	path string

	files []repository.File
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
	files, err := s.fs.Glob(fmt.Sprintf("%s/*.meta4", s.path))
	if err != nil {
		return bosherr.WrapError(err, "Listing metalinks")
	}

	s.files = []repository.File{}

	for _, file := range files {
		stat, err := s.fs.Stat(file)
		if err != nil {
			return bosherr.WrapError(err, "Stat receipt")
		}

		metalinkBytes, err := s.fs.ReadFile(file)
		if err != nil {
			return bosherr.WrapError(err, "Reading receipt")
		}

		results, err := source.ExplodeMetalinkBytes(
			repository.Repository{
				URI:     s.URI(),
				Path:    path.Base(file),
				Version: stat.ModTime().Format(time.RFC3339),
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
	return s.uri
}

func (s Source) FilterFiles(filter filter.Filter) ([]repository.File, error) {
	return source.FilterFilesInMemory(s.files, filter)
}
