package cmd

import (
	"fmt"

	"github.com/dpb587/metalink/repository/cli/cmd/args"
	"github.com/dpb587/metalink/repository/filter"
	filter_and "github.com/dpb587/metalink/repository/filter/and"
	"github.com/dpb587/metalink/repository/formatter"
	"github.com/dpb587/metalink/repository/sorter"
	sorter_fileversion "github.com/dpb587/metalink/repository/sorter/fileversion"
	sorter_reverse "github.com/dpb587/metalink/repository/sorter/reverse"
	// "github.com/dpb587/metalink/repository/sorter"
	"github.com/dpb587/metalink/repository"
	"github.com/dpb587/metalink/repository/source"

	bosherr "github.com/cloudfoundry/bosh-utils/errors"
)

type Filter struct {
	Filter      []args.Filter `short:"f" long:"filter" description:"Filter metalink files" default-mask:"TYPE[:VALUE]"`
	SortReverse bool          `long:"sort-reverse" description:"Reverse sort order"`
	Limit       int           `short:"n" long:"limit" description:"Limit the number of metalink files"`
	Args        FilterArgs    `positional-args:"true" required:"true"`
	Format      string        `long:"format" description:"Format to dump the repository (json, version, xml)" default:"xml"`

	SourceFactory source.Factory
	FilterManager filter.Manager
}

type FilterArgs struct {
	RepositoryURI string `positional-arg-name:"URI" description:"Repository URI hosting the files"`
}

func (c *Filter) Execute(_ []string) error {
	source, err := c.SourceFactory.Create(c.Args.RepositoryURI)
	if err != nil {
		return bosherr.WrapError(err, "Creating repository")
	}

	err = source.Load()
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

	metalinks, err := source.Filter(andFilter)
	if err != nil {
		return bosherr.WrapError(err, "Filtering metalink files")
	}

	var sort sorter.Sorter = sorter_fileversion.Sorter{}

	if c.SortReverse {
		sort = sorter_reverse.Sorter{Sorter: sort}
	}

	sorter.Sort(metalinks, sort)

	limit := c.Limit
	if limit < 0 {
		panic("Invalid limit")
	} else if limit == 0 {
		limit = 10
	}

	filtered := []repository.RepositoryMetalink{}

	for _, meta4 := range metalinks {
		filtered = append(filtered, meta4)

		if limit--; limit <= 0 {
			break
		}
	}

	switch c.Format {
	case "json":
		return formatter.JSONFormatter{}.DumpRepository(filtered)
	case "version":
		return formatter.VersionFormatter{}.DumpRepository(filtered)
	case "xml":
		return formatter.XMLFormatter{}.DumpRepository(filtered)
	default:
		return fmt.Errorf("Unknown format: %s", c.Format)
	}
}
