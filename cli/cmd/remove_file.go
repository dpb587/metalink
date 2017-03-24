package cmd

import "errors"

type RemoveFile struct {
	Meta4
	Args RemoveFileArgs `positional-args:"true" required:"true"`
}

type RemoveFileArgs struct {
	Name string `positional-arg-name:"NAME" description:"File name"`
}

func (c *RemoveFile) Execute(_ []string) error {
	meta4, err := c.Meta4.Get()
	if err != nil {
		return err
	}

	for fileIdx, file := range meta4.Files {
		if file.Name != c.Args.Name {
			continue
		}

		meta4.Files = append(meta4.Files[:fileIdx], meta4.Files[fileIdx+1:]...)

		return c.Meta4.Put(meta4)
	}

	return errors.New("File does not exist")
}
