package cmd

import "errors"

type FileRemoveURL struct {
	Meta4File
	Args FileRemoveURLArgs `positional-args:"true" required:"true"`
}

type FileRemoveURLArgs struct {
	URL string `positional-arg-name:"URL" description:"Download URI"`
}

func (c *FileRemoveURL) Execute(_ []string) error {
	file, err := c.Meta4File.Get()
	if err != nil {
		return err
	}

	for urlIdx, url := range file.URLs {
		if url.URL != c.Args.URL {
			continue
		}

		file.URLs = append(file.URLs[:urlIdx], file.URLs[urlIdx+1:]...)

		return c.Meta4File.Put(file)
	}

	return errors.New("File does not exist")
}
