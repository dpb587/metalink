package cmd

import "github.com/dpb587/blob-receipt"

type FileSetHash struct {
	Meta4File
	Args FileSetHashArgs `positional-args:"true" required:"true"`
}

type FileSetHashArgs struct {
	Type string `positional-arg-name:"TYPE" description:"Hash algorithm (md5 sha-256 sha-512)"`
	Hash string `positional-arg-name:"HASH" description:"Hash"`
}

func (c *FileSetHash) Execute(_ []string) error {
	file, err := c.Meta4File.Get()
	if err != nil {
		return err
	}

	for hashIdx, hash := range file.Hashes {
		if hash.Type == c.Args.Type {
			file.Hashes = append(file.Hashes[:hashIdx], file.Hashes[hashIdx+1:]...)

			break
		}
	}

	file.Hashes = append(
		file.Hashes,
		blobreceipt.Hash{
			Type: c.Args.Type,
			Hash: c.Args.Hash,
		},
	)

	return c.Meta4File.Put(file)
}
