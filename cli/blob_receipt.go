package main

import (
	"os"

	"github.com/dpb587/blob-receipt/cli/cmd"
	"github.com/dpb587/blob-receipt/origin"
	"github.com/dpb587/blob-receipt/storage"
	flags "github.com/jessevdk/go-flags"

	boshlog "github.com/cloudfoundry/bosh-utils/logger"
	boshsys "github.com/cloudfoundry/bosh-utils/system"
)

func main() {
	logger := boshlog.NewLogger(boshlog.LevelError)
	fs := boshsys.NewOsFileSystem(logger)

	originFactory := origin.NewDefaultFactory(fs)
	storageFactory := storage.NewDefaultFactory(fs)

	c := struct {
		Create cmd.Create `command:"create" description:"Create or update a receipt for a given blob"`
		Verify cmd.Verify `command:"verify" description:"Verify size and digest of a receipt match a given blob"`
		// Download cmd.Download `command:"download" description:"Download (and verify) from a receipt to a local file"`
	}{
		Create: cmd.Create{
			OriginFactory:  originFactory,
			StorageFactory: storageFactory,
		},
		Verify: cmd.Verify{
			OriginFactory:  originFactory,
			StorageFactory: storageFactory,
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
