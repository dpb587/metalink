package template

import (
	"github.com/dpb587/metalink"
)

type templateFile metalink.File

func (tf templateFile) MD5() string {
	for _, hash := range tf.Hashes {
		if hash.Type == "md5" {
			return hash.Hash
		}
	}

	return ""
}

func (tf templateFile) SHA1() string {
	for _, hash := range tf.Hashes {
		if hash.Type == "sha-1" {
			return hash.Hash
		}
	}

	return ""
}

func (tf templateFile) SHA256() string {
	for _, hash := range tf.Hashes {
		if hash.Type == "sha-256" {
			return hash.Hash
		}
	}

	return ""
}

func (tf templateFile) SHA512() string {
	for _, hash := range tf.Hashes {
		if hash.Type == "sha-512" {
			return hash.Hash
		}
	}

	return ""
}
