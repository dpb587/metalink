package cmd

import (
	"fmt"

	"github.com/dpb587/metalink/repository/cli/cmd/args"
	"github.com/dpb587/metalink/repository/filter"
	filter_and "github.com/dpb587/metalink/repository/filter/and"
	"github.com/dpb587/metalink/repository/sorter"
	sorter_fileversion "github.com/dpb587/metalink/repository/sorter/fileversion"
	sorter_reverse "github.com/dpb587/metalink/repository/sorter/reverse"
	// "github.com/dpb587/metalink/repository/sorter"
	"github.com/dpb587/metalink/repository/source"

	bosherr "github.com/cloudfoundry/bosh-utils/errors"
)

type Versions struct {
	Filter      []args.Filter `short:"f" long:"filter" description:"Filter metalink files" default-mask:"TYPE[:VALUE]"`
	SortReverse bool          `long:"sort-reverse" description:"Reverse sort order"`
	Limit       int           `short:"n" long:"limit" description:"Limit the number of metalink files"`
	Args        VersionsArgs  `positional-args:"true" required:"true"`

	SourceFactory source.Factory
	FilterManager filter.Manager
}

type VersionsArgs struct {
	RepositoryURI string `positional-arg-name:"URI" description:"Repository URI hosting the files"`
}

func (c *Versions) Execute(_ []string) error {
	repository, err := c.SourceFactory.Create(c.Args.RepositoryURI)
	if err != nil {
		return bosherr.WrapError(err, "Creating repository")
	}

	err = repository.Load()
	if err != nil {
		return bosherr.WrapError(err, "Loading repository")
	}

	andFilter := filter_and.NewFilter()

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

	var sort sorter.Sorter = sorter_fileversion.Sorter{}

	if c.SortReverse {
		sort = sorter_reverse.Sorter{Sorter: sort}
	}

	sorter.Sort(files, sort)

	limit := c.Limit
	if limit < 0 {
		panic("Invalid limit")
	} else if limit == 0 {
		limit = 10
	}

	versionsSeen := map[string]bool{}

	for _, file := range files {
		if _, seen := versionsSeen[file.File.Version]; seen {
			continue
		}

		fmt.Println(file.File.Version)

		versionsSeen[file.File.Version] = true

		if limit--; limit <= 0 {
			break
		}
	}

	return nil
}
