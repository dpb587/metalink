package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/dpb587/metalink"
	"github.com/dpb587/metalink/repository/cli/cmd/args"
	"github.com/dpb587/metalink/repository/filter"
	filter_and "github.com/dpb587/metalink/repository/filter/and"
	filter_fileversion "github.com/dpb587/metalink/repository/filter/fileversion"
	"github.com/dpb587/metalink/storage"
	// "github.com/dpb587/metalink/repository/sorter"
	"github.com/dpb587/metalink/repository/source"

	bosherr "github.com/cloudfoundry/bosh-utils/errors"
)

type Version struct {
	Filter []args.Filter `short:"f" long:"filter" description:"Filter metalink files" default-mask:"TYPE[:VALUE]"`
	Args   VersionArgs   `positional-args:"true" required:"true"`

	SourceFactory source.Factory
	FilterManager filter.Manager
}

type VersionArgs struct {
	RepositoryURI string `positional-arg-name:"URI" description:"Repository URI hosting the files"`
	Version       string `positionl-arg-name:"VERSION" description:"Semantic Version"`
}

func (c *Version) Execute(_ []string) error {
	repository, err := c.SourceFactory.Create(c.Args.RepositoryURI)
	if err != nil {
		return bosherr.WrapError(err, "Creating repository")
	}

	err = repository.Reload()
	if err != nil {
		return bosherr.WrapError(err, "Loading repository")
	}

	andFilter := filter_and.NewFilter()

	versionFilter, err := filter_fileversion.CreateFilter(c.Args.Version)
	if err != nil {
		return bosherr.WrapError(err, "Parsing version")
	}

	andFilter.Add(versionFilter)

	for filterArgIdx, filterArg := range c.Filter {
		addFilter, err := c.FilterManager.CreateFilter(filterArg.Type, filterArg.Value)
		if err != nil {
			return bosherr.WrapErrorf(err, "Parsing filter argument %d", filterArgIdx)
		}

		andFilter.Add(addFilter)
	}

	files, err := repository.FilterFiles(andFilter)
	if err != nil {
		return bosherr.WrapError(err, "Filtering metalink files")
	}

	dynamic := true
	meta4 := metalink.Metalink{
		Generator: "repository.metalink.dpb587.github.io/0.0.0",
		Origin: &metalink.Origin{
			Dynamic: &dynamic,
			URL: fmt.Sprintf(
				"metalink-repository:version?repository_uri=%s&version=%s",
				url.QueryEscape(c.Args.RepositoryURI),
				url.QueryEscape(c.Args.Version),
			),
		},
		Files: []metalink.File{},
	}

	for _, file := range files {
		meta4.Files = append(meta4.Files, file.File)
	}

	return storage.WriteMetalink(os.Stdout, meta4)
}
