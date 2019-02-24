package cmd

import (
	"os"
	"path/filepath"

	"github.com/dpb587/metalink/repository/filter"
	"github.com/dpb587/metalink/repository/source"

	"github.com/pkg/errors"
)

type Put struct {
	Args PutArgs `positional-args:"true" required:"true"`

	SourceFactory source.Factory
	FilterManager filter.Manager
}

type PutArgs struct {
	RepositoryURI string `positional-arg-name:"URI" description:"Repository URI hosting the files"`
	MetalinkPath  string `positional-arg-name:"METALINK-PATH" description:"Metalink path"`
}

func (c *Put) Execute(_ []string) error {
	repository, err := c.SourceFactory.Create(c.Args.RepositoryURI, map[string]interface{}{})
	if err != nil {
		return errors.Wrap(err, "Creating repository")
	}

	fh, err := os.Open(c.Args.MetalinkPath)
	if err != nil {
		return errors.Wrap(err, "Opening metalink")
	}

	defer fh.Close()

	err = repository.Put(filepath.Base(c.Args.MetalinkPath), fh)
	if err != nil {
		return errors.Wrap(err, "Putting metalink")
	}

	return nil
}
