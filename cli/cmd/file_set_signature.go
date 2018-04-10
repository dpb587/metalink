package cmd

import (
	"io/ioutil"
	"os"

	"github.com/pkg/errors"
	boshsys "github.com/cloudfoundry/bosh-utils/system"
	"github.com/dpb587/metalink"
)

type FileSetSignature struct {
	Meta4File
	FS   boshsys.FileSystem
	Args FileSetSignatureArgs `positional-args:"true" required:"true"`
}

type FileSetSignatureArgs struct {
	MediaType string `positional-arg-name:"MEDIA-TYPE" description:"Signature media type"`
	File      string `positional-arg-name:"FILE" description:"Path to file with signature contents"`
}

func (c *FileSetSignature) Execute(_ []string) error {
	file, err := c.Meta4File.Get()
	if err != nil {
		return err
	}

	var signature []byte

	if c.Args.File == "-" {
		signature, err = ioutil.ReadAll(os.Stdin)
		if err != nil {
			return errors.Wrap(err, "Reading stdin")
		}
	} else {
		signature, err = c.FS.ReadFile(c.Args.File)
		if err != nil {
			return errors.Wrap(err, "Reading file")
		}
	}

	file.Signature = &metalink.Signature{
		MediaType: c.Args.MediaType,
		Signature: string(signature),
	}

	return c.Meta4File.Put(file)
}
