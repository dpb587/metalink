package hash

import (
	"fmt"

	"github.com/dpb587/metalink"
	"github.com/dpb587/metalink/verification"
)

type Result struct {
	hash metalink.Hash
}

var _ verification.Result = Result{}

func NewResult(hash metalink.Hash) Result {
	return Result{
		hash: hash,
	}
}

func (v Result) Apply(meta4 *metalink.File) error {
	if _, found := Find(*meta4, v.hash.Type); found {
		return fmt.Errorf("hash already exists: %s", v.hash.Type)
	}

	meta4.Hashes = append(meta4.Hashes, v.hash)

	return nil
}

func (v Result) Verify(meta4 metalink.File) error {
	expected, found := Find(meta4, v.hash.Type)
	if !found {
		return fmt.Errorf("hash not found: %s", v.hash.Type)
	}

	if v.hash.Hash != expected.Hash {
		return fmt.Errorf("expected hash to be %s", expected.Hash)
	}

	return nil
}

func (v Result) Type() string {
	return v.hash.Type
}

func (v Result) Summary() string {
	return v.hash.Hash
}
