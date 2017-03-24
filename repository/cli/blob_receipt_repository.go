package main

import (
	"os"

	"github.com/dpb587/metalink/repository/cli/cmd"
	"github.com/dpb587/metalink/repository/filterfactory"
	source_factory "github.com/dpb587/metalink/repository/source/factory"
	source_fs "github.com/dpb587/metalink/repository/source/fs"
	source_git "github.com/dpb587/metalink/repository/source/git"

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
	sourceFactory.Add(source_git.NewFactory(fs, cmdRunner))

	filterManager := filterfactory.NewManager()

	c := struct {
		List cmd.List `command:"list" description:"List blob receipts in a repository"`
	}{
		List: cmd.List{
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
