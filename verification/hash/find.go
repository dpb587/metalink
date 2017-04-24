package hash

import "github.com/dpb587/metalink"

func Find(meta4 metalink.File, hashType string) (metalink.Hash, bool) {
	for _, found := range meta4.Hashes {
		if found.Type == hashType {
			return found, true
		}
	}

	return metalink.Hash{}, false
}
