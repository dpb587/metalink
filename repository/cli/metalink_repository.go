package main

import (
	"os"

	"github.com/dpb587/metalink/repository/cli/cmd"
	"github.com/dpb587/metalink/repository/filterfactory"
	source_factory "github.com/dpb587/metalink/repository/source/factory"
	source_fs "github.com/dpb587/metalink/repository/source/fs"
	source_git "github.com/dpb587/metalink/repository/source/git"
	source_http "github.com/dpb587/metalink/repository/source/http"
	source_s3 "github.com/dpb587/metalink/repository/source/s3"

	flags "github.com/jessevdk/go-flags"

	boshlog "github.com/cloudfoundry/bosh-utils/logger"
	boshsys "github.com/cloudfoundry/bosh-utils/system"
)

func main() {
	logger := boshlog.NewLogger(boshlog.LevelError)
	fs := boshsys.NewOsFileSystem(logger)
	cmdRunner := boshsys.NewExecCmdRunner(logger)

	sourceFactory := source_factory.NewFactory()
	sourceFactory.Add(source_fs.NewFactory(fs))
	sourceFactory.Add(source_http.NewFactory())
	sourceFactory.Add(source_git.NewFactory(fs, cmdRunner))
	sourceFactory.Add(source_s3.NewFactory())

	filterManager := filterfactory.NewManager()

	c := struct {
		Filter cmd.Filter `command:"filter" description:"Filter metalinks from a repository"`
		Show   cmd.Show   `command:"show" description:"Show a metalink from a repository"`
		Put    cmd.Put    `command:"put" description:"Put a metalink to a repository"`
	}{
		Filter: cmd.Filter{
			SourceFactory: sourceFactory,
			FilterManager: filterManager,
		},
		Show: cmd.Show{
			SourceFactory: sourceFactory,
			FilterManager: filterManager,
		},
		Put: cmd.Put{
			SourceFactory: sourceFactory,
			FilterManager: filterManager,
		},
	}

	var parser = flags.NewParser(&c, flags.Default)
	if _, err := parser.Parse(); err != nil {
		if flagsErr, ok := err.(*flags.Error); ok && flagsErr.Type == flags.ErrHelp {
			os.Exit(0)
		} else {
			os.Exit(1)
		}
	}
}
