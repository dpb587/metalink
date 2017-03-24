package cmd

import "fmt"

type FileURLs struct {
	Meta4File
}

func (c *FileURLs) Execute(_ []string) error {
	file, err := c.Meta4File.Get()
	if err != nil {
		return err
	}

	for _, url := range file.URLs {
		fmt.Println(url.URL)
	}

	return nil
}
