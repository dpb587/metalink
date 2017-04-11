package cmd

import (
	"fmt"

	"github.com/dpb587/metalink/repository/cli/cmd/args"
	"github.com/dpb587/metalink/repository/filter"
	filter_and "github.com/dpb587/metalink/repository/filter/and"
	"github.com/dpb587/metalink/repository/formatter"
	"github.com/dpb587/metalink/repository/source"

	bosherr "github.com/cloudfoundry/bosh-utils/errors"
)

type Show struct {
	Filter []args.Filter `short:"f" long:"filter" description:"Filter metalink files" default-mask:"TYPE[:VALUE]"`
	Args   ShowArgs      `positional-args:"true" required:"true"`
	Format string        `long:"format" description:"Format to dump the metalink (json, version, xml)" default:"xml"`

	SourceFactory source.Factory
	FilterManager filter.Manager
}

type ShowArgs struct {
	RepositoryURI string `positional-arg-name:"URI" description:"Repository URI hosting the files"`
}

func (c *Show) Execute(_ []string) error {
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

	metalinks, err := repository.Filter(andFilter)
	if err != nil {
		return bosherr.WrapError(err, "Filtering metalinks")
	}

	if len(metalinks) != 1 {
		return fmt.Errorf("Expected 1 metalink, but found %d", len(metalinks))
	}

	switch c.Format {
	case "json":
		return formatter.JSONFormatter{}.DumpMetalink(metalinks[0])
	case "version":
		return formatter.VersionFormatter{}.DumpMetalink(metalinks[0])
	case "xml":
		return formatter.XMLFormatter{}.DumpMetalink(metalinks[0])
	default:
		return fmt.Errorf("Unknown format: %s", c.Format)
	}
}
