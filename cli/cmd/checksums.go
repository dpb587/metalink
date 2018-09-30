package cmd

import (
	"fmt"

	"github.com/dpb587/metalink"
)

type Checksums struct {
	Meta4

	Type metalink.HashType `long:"type" description:"Hash type to show" default:"sha-256"`
}

func (c *Checksums) Execute(_ []string) error {
	meta4, err := c.Meta4.Get()
	if err != nil {
		return err
	}

	for _, file := range meta4.Files {
		for _, hash := range file.Hashes {
			if hash.Type != c.Type {
				continue
			}

			fmt.Printf("%s  %s\n", hash.Hash, file.Name)

			break
		}
	}

	return nil
}
