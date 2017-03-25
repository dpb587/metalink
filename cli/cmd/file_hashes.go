package cmd

import "fmt"

type FileHashes struct {
	Meta4File
}

func (c *FileHashes) Execute(_ []string) error {
	file, err := c.Meta4File.Get()
	if err != nil {
		return err
	}

	for _, hash := range file.Hashes {
		fmt.Println(fmt.Sprintf("%s %s", hash.Type, hash.Hash))
	}

	return nil
}
