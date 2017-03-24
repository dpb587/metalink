package cmd

import (
	"errors"

	blobreceipt "github.com/dpb587/blob-receipt"
)

type AddFile struct {
	Meta4
	Args AddFileArgs `positional-args:"true" required:"true"`
}

type AddFileArgs struct {
	Name string `positional-arg-name:"NAME" description:"File name"`
}

func (c *AddFile) Execute(_ []string) error {
	meta4, err := c.Meta4.Get()
	if err != nil {
		return err
	}

	for _, file := range meta4.Files {
		if file.Name != c.Args.Name {
			continue
		}

		return errors.New("File already exists")
	}

	meta4.Files = append(
		meta4.Files,
		blobreceipt.File{
			Name: c.Args.Name,
		},
	)

	return c.Meta4.Put(meta4)
}
