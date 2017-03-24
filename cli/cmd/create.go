package cmd

import (
	"errors"

	"github.com/dpb587/metalink"
)

type Create struct {
	Meta4
}

func (c *Create) Execute(_ []string) error {
	exists, err := c.Meta4.Exists()
	if err != nil {
		return err
	} else if exists {
		return errors.New("Metalink file already exists")
	}

	return c.Meta4.Put(metalink.Metalink{Generator: "metalink.dpb587.github.io/0.0.0"})
}
