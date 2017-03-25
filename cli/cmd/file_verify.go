package cmd

import (
	"errors"
	"fmt"
	"strings"

	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	"github.com/dpb587/metalink"
	"github.com/dpb587/metalink/crypto"
	"github.com/dpb587/metalink/origin"
)

type FileVerify struct {
	Meta4File
	OriginFactory origin.OriginFactory `no-flag:"true"`
	Quiet         bool                 `long:"quiet" short:"q" description:"Suppress passing digests"`
	Hashes        []string             `long:"hash" description:"Specific hash type(s) to verify; or 'all'" default-mask:"strongest available"`
	Args          FileVerifyArgs       `positional-args:"true" required:"true"`
}

type FileVerifyArgs struct {
	Local string `positional-arg-name:"PATH" description:"Path to the blob file"`
}

func (c *FileVerify) Execute(_ []string) error {
	file, err := c.Meta4File.Get()
	if err != nil {
		return err
	}

	local, err := c.OriginFactory.New(c.Args.Local)
	if err != nil {
		return bosherr.WrapError(err, "Parsing origin destination")
	}

	verifyHashes := c.Hashes
	knownHashes := []string{}

	for _, hash := range file.Hashes {
		knownHashes = append(knownHashes, hash.Type)
	}

	if len(c.Hashes) == 0 {
		algorithm, _ := crypto.GetStrongestAlgorithm(knownHashes)

		verifyHashes = []string{crypto.GetDigestType(algorithm.Name())}
	} else if len(c.Hashes) == 1 && c.Hashes[0] == "all" {
		verifyHashes = knownHashes
	}

	if len(verifyHashes) == 0 {
		return errors.New("Failed to find a hash")
	}

	failed := []string{}

	for _, verifyHashType := range verifyHashes {
		var expectedHash *metalink.Hash

		for _, seekHash := range file.Hashes {
			if seekHash.Type != verifyHashType {
				continue
			}

			expectedHash = &seekHash

			break
		}

		if expectedHash == nil {
			failed = append(failed, verifyHashType)

			continue
		}

		reader, err := local.Reader()
		if err != nil {
			return bosherr.WrapErrorf(err, "Opening origin for %s", verifyHashType)
		}

		algorithm, err := crypto.GetAlgorithm(expectedHash.Type)
		if err != nil {
			fmt.Println(fmt.Sprintf("FAIL\t%s\tERROR\t%s", verifyHashType, err))

			failed = append(failed, verifyHashType)

			continue
		}

		actualHash, err := algorithm.CreateDigest(reader)
		if err != nil {
			fmt.Println(fmt.Sprintf("FAIL\t%s\tERROR\t%s", verifyHashType, err))

			failed = append(failed, verifyHashType)

			continue
		}

		if expectedHash.Hash == crypto.GetDigestHash(actualHash) {
			if !c.Quiet {
				fmt.Println(fmt.Sprintf("OKAY\t%s\t%s", expectedHash.Type, expectedHash.Hash))
			}

			continue
		}

		fmt.Println(fmt.Sprintf("FAIL\t%s\t%s\t%s", expectedHash.Type, expectedHash.Hash, crypto.GetDigestHash(actualHash)))

		failed = append(failed, verifyHashType)
	}

	if len(failed) > 0 {
		return fmt.Errorf("Failed to verify digest: %s", strings.Join(failed, ", "))
	}

	return nil
}
