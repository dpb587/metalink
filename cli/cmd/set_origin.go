package cmd

import "github.com/dpb587/metalink"

type SetOrigin struct {
	Meta4
	Dynamic *bool         `long:"dynamic" description:"If origin contains dynamic updated information"`
	Args    SetOriginArgs `positional-args:"true" required:"true"`
}

type SetOriginArgs struct {
	URL string `positional-arg-name:"URL" description:"Origin IRI"`
}

func (c *SetOrigin) Execute(_ []string) error {
	meta4, err := c.Meta4.Get()
	if err != nil {
		return err
	}

	meta4.Origin = &metalink.Origin{
		Dynamic: c.Dynamic,
		URL:     c.Args.URL,
	}

	return c.Meta4.Put(meta4)
}
