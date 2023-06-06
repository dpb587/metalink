package cmd

import (
	"fmt"

	"github.com/dpb587/metalink/storage"
	"github.com/pkg/errors"
)

type ImportMetalink struct {
	Meta4
	StorageFactory storage.StorageFactory `no-flag:"true"`

	Merge bool           `long:"merge" description:"If existing file, overwrite fields"`
	Args  ImportFileArgs `positional-args:"true" required:"true"`
}

type ImportMetalinkArgs struct {
	Path string `positional-arg-name:"PATH" description:"Metalink path"`
}

func (c *ImportMetalink) Execute(_ []string) error {
	meta4, err := c.Meta4.Get()
	if err != nil {
		return err
	}

	storageImport, err := c.StorageFactory.New(c.Args.Path)
	if err != nil {
		return errors.Wrap(err, "Preparing import storage")
	}

	meta4Import, err := storageImport.Get()
	if err != nil {
		return errors.Wrap(err, "Loading metalink import")
	}

	for _, file := range meta4Import.Files {
		for fileIdx, existingFile := range meta4.Files {
			if existingFile.Name != file.Name {
				continue
			}

			if !c.Merge {
				return fmt.Errorf("File already exists: %s", existingFile.Name)
			}

			meta4.Files = append(meta4.Files[:fileIdx], meta4.Files[fileIdx+1:]...)

			break
		}

		meta4.Files = append(meta4.Files, file)
	}

	return c.Meta4.Put(meta4)
}
