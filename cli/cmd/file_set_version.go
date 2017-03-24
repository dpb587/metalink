package cmd

type FileSetVersion struct {
	Meta4File
	Args FileSetVersionArgs `positional-args:"true" required:"true"`
}

type FileSetVersionArgs struct {
	Version string `positional-arg-name:"VERSION" description:"File version"`
}

func (c *FileSetVersion) Execute(_ []string) error {
	file, err := c.Meta4File.Get()
	if err != nil {
		return err
	}

	file.Version = c.Args.Version

	return c.Meta4File.Put(file)
}
