package main

import (
	"os"

	"github.com/dpb587/metalink/repository/cli/cmd"
	"github.com/dpb587/metalink/repository/filterfactory"
	source_factory "github.com/dpb587/metalink/repository/source/factory"
	source_fs "github.com/dpb587/metalink/repository/source/fs"
	source_git "github.com/dpb587/metalink/repository/source/git"
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
	sourceFactory.Add(source_git.NewFactory(fs, cmdRunner))
	sourceFactory.Add(source_s3.NewFactory(fs, cmdRunner))

	filterManager := filterfactory.NewManager()

	c := struct {
		Versions cmd.Versions `command:"versions" description:"Versions in a repository"`
		Version  cmd.Version  `command:"version" description:"Get metalink for a specific version in a repository"`
	}{
		Versions: cmd.Versions{
			SourceFactory: sourceFactory,
			FilterManager: filterManager,
		},
		Version: cmd.Version{
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
