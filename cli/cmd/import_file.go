package cmd

import (
	"errors"
	"fmt"
	"path"

	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	"github.com/dpb587/metalink"
	"github.com/dpb587/metalink/file/url"
	"github.com/dpb587/metalink/verification"
	"github.com/dpb587/metalink/verification/hash"
)

type ImportFile struct {
	Meta4File
	URLLoader url.Loader

	Merge   bool           `long:"merge" description:"If existing file, overwrite fields"`
	Hashes  []string       `long:"hash" description:"Specific hashes to calculate" default:"md5" default:"sha-1" default:"sha-256" default:"sha-512"`
	Version string         `long:"version" description:"File version"`
	Args    ImportFileArgs `positional-args:"true" required:"true"`
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

	origin, err := c.URLLoader.Load(metalink.URL{URL: c.Args.Path})
	if err != nil {
		return bosherr.WrapError(err, "Loading origin")
	}

	file.Size, err = origin.Size()
	if err != nil {
		return bosherr.WrapError(err, "Loading size")
	}

	hashmap := map[string]verification.Signer{
		"sha-512": hash.SHA512Verification,
		"sha-256": hash.SHA256Verification,
		"sha-1":   hash.SHA1Verification,
		"md5":     hash.MD5Verification,
	}

	for _, hashType := range c.Hashes {
		signer, found := hashmap[hashType]
		if !found {
			return fmt.Errorf("unknown hash type: %s", hashType)
		}

		verification, err := signer.Sign(origin)
		if err != nil {
			return bosherr.WrapError(err, "Signing hash")
		}

		err = verification.Apply(&file)
		if err != nil {
			return bosherr.WrapError(err, "Adding verification to file")
		}
	}

	meta4.Files = append(meta4.Files, file)

	return c.Meta4.Put(meta4)
}
