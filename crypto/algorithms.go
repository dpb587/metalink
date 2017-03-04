package crypto

import (
	"fmt"
	"strings"

	boshcry "github.com/cloudfoundry/bosh-utils/crypto"
)

var algorithmStrength = []string{
	"sha512",
	"sha256",
	"sha1",
	"md5",
}

func GetAlgorithm(algorithm string) (boshcry.Algorithm, error) {
	switch algorithm {
	case "md5":
		return DigestAlgorithmMD5, nil
	case "sha1":
		return boshcry.DigestAlgorithmSHA1, nil
	case "sha256":
		return boshcry.DigestAlgorithmSHA256, nil
	case "sha512":
		return boshcry.DigestAlgorithmSHA512, nil
	}

	return nil, fmt.Errorf("Unsupported digest algorithm: %s", algorithm)
}

func GetStrongestAlgorithm(algorithms []string) (boshcry.Algorithm, error) {
	for _, candidateAlgorithm := range algorithmStrength {
		for _, algorithm := range algorithms {
			if algorithm != candidateAlgorithm {
				continue
			}

			return GetAlgorithm(algorithm)
		}
	}

	return nil, fmt.Errorf("No strong algorithm found: %s", strings.Join(algorithms, ", "))
}

func GetDigestHash(digest boshcry.Digest) string {
	digestHash := digest.String()

	if digest.Algorithm().Name() == "sha1" {
		return digestHash
	}

	digestHashParts := strings.SplitN(digestHash, ":", 2)
	if len(digestHashParts) != 2 {
		panic(fmt.Sprintf("expected hash in format of 'algorithm:hash' but received '%s'", digestHash))
	}

	return digestHashParts[1]
}
