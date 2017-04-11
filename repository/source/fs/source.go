package fs

import (
	"fmt"
	"path"
	"time"

	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	boshsys "github.com/cloudfoundry/bosh-utils/system"
	"github.com/dpb587/metalink"
	"github.com/dpb587/metalink/repository"
	"github.com/dpb587/metalink/repository/filter"
	"github.com/dpb587/metalink/repository/source"
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

	files, err := s.fs.Glob(fmt.Sprintf("%s/*.meta4", s.path))
	if err != nil {
		return bosherr.WrapError(err, "Listing metalinks")
	}

	for _, file := range files {
		stat, err := s.fs.Stat(file)
		if err != nil {
			return bosherr.WrapError(err, "Stat receipt")
		}

		metalinkBytes, err := s.fs.ReadFile(file)
		if err != nil {
			return bosherr.WrapError(err, "Reading metalink")
		}

		repometa4 := repository.RepositoryMetalink{
			Reference: repository.RepositoryMetalinkReference{
				Repository: uri,
				Path:       path.Base(file),
				Version:    stat.ModTime().Format(time.RFC3339),
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
	return s.uri
}

func (s Source) Filter(f filter.Filter) ([]repository.RepositoryMetalink, error) {
	return source.FilterInMemory(s.metalinks, f)
}
