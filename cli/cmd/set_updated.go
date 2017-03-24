package cmd

import "github.com/dpb587/metalink/cli/args"

type SetUpdated struct {
	Meta4
	Args SetUpdatedArgs `positional-args:"true" required:"true"`
}

type SetUpdatedArgs struct {
	Time args.Time `positional-arg-name:"TIME" description:"A date/time"`
}

func (c *SetUpdated) Execute(_ []string) error {
	meta4, err := c.Meta4.Get()
	if err != nil {
		return err
	}

	meta4.Updated = &c.Args.Time.Time

	return c.Meta4.Put(meta4)
}
