package cmd

import (
	"errors"
	"fmt"

	"github.com/dpb587/metalink"
)

type FileHash struct {
	Meta4File
	Args FileHashArgs `positional-args:"true" required:"true"`
}

type FileHashArgs struct {
	Type metalink.HashType `positional-arg-name:"TYPE" description:"Hash algorithm (md5 sha-256 sha-512)"`
}

func (c *FileHash) Execute(_ []string) error {
	file, err := c.Meta4File.Get()
	if err != nil {
		return err
	}

	for _, hash := range file.Hashes {
		if hash.Type != c.Args.Type {
			continue
		}

		fmt.Println(hash.Hash)

		return nil
	}

	return errors.New("Hash does not exist")
}
