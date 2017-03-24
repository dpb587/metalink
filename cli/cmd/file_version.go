package cmd

import (
	"errors"
	"fmt"
)

type FileVersion struct {
	Meta4File
}

func (c *FileVersion) Execute(_ []string) error {
	file, err := c.Meta4File.Get()
	if err != nil {
		return err
	}

	if file.Version == "" {
		return errors.New("Version is not set")
	}

	fmt.Println(file.Version)

	return nil
}
