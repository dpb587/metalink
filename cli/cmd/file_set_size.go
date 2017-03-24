package cmd

type FileSetSize struct {
	Meta4File
	Args FileSetSizeArgs `positional-args:"true" required:"true"`
}

type FileSetSizeArgs struct {
	Size uint64 `positional-arg-name:"SIZE" description:"File size"`
}

func (c *FileSetSize) Execute(_ []string) error {
	file, err := c.Meta4File.Get()
	if err != nil {
		return err
	}

	file.Size = c.Args.Size

	return c.Meta4File.Put(file)
}
