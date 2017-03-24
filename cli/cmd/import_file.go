package cmd

import (
	"errors"
	"path"

	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	"github.com/dpb587/metalink"
	"github.com/dpb587/metalink/crypto"
	"github.com/dpb587/metalink/origin"
)

type ImportFile struct {
	Meta4File
	OriginFactory origin.OriginFactory
	Merge         bool           `long:"merge" description:"If existing file, overwrite fields"`
	Hashes        []string       `long:"hash" description:"Specific hashes to calculate" default:"md5" default:"sha-1" default:"sha-256" default:"sha-512"`
	Version       string         `long:"version" description:"File version"`
	Args          ImportFileArgs `positional-args:"true" required:"true"`
}

type ImportFileArgs struct {
	Path string `positional-arg-name:"PATH" description:"File path"`
}

func (c *ImportFile) Execute(_ []string) error {
	meta4, err := c.Meta4.Get()
	if err != nil {
		return err
	}

	fileName := c.Meta4File.File
	if fileName == "" {
		fileName = path.Base(c.Args.Path)
	}

	file := metalink.File{
		Name:    fileName,
		Version: c.Version,
		Hashes:  []metalink.Hash{},
	}

	for fileIdx, existingFile := range meta4.Files {
		if existingFile.Name != fileName {
			continue
		}

		if !c.Merge {
			return errors.New("File already exists")
		}

		meta4.Files = append(meta4.Files[:fileIdx], meta4.Files[fileIdx+1:]...)

		file = existingFile

		break
	}

	origin, err := c.OriginFactory.New(c.Args.Path)
	if err != nil {
		return bosherr.WrapError(err, "Loading origin")
	}

	file.Size, err = origin.Size()
	if err != nil {
		return bosherr.WrapError(err, "Loading size")
	}

	for _, algorithmName := range c.Hashes {
		algorithm, err := crypto.GetAlgorithm(algorithmName)
		if err != nil {
			return bosherr.WrapErrorf(err, "Loading digest algorithm")
		}

		digest, err := origin.Digest(algorithm)
		if err != nil {
			return bosherr.WrapErrorf(err, "Sourcing blob %s digest", algorithm.Name())
		}

		file.Hashes = append(
			file.Hashes,
			metalink.Hash{
				Type: crypto.GetDigestType(algorithm.Name()),
				Hash: crypto.GetDigestHash(digest),
			},
		)
	}

	meta4.Files = append(meta4.Files, file)

	return c.Meta4.Put(meta4)
}
