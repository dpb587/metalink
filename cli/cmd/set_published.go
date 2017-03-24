package cmd

import "github.com/dpb587/blob-receipt/cli/args"

type SetPublished struct {
	Meta4
	Args SetPublishedArgs `positional-args:"true" required:"true"`
}

type SetPublishedArgs struct {
	Time args.Time `positional-arg-name:"TIME" description:"A date/time"`
}

func (c *SetPublished) Execute(_ []string) error {
	meta4, err := c.Meta4.Get()
	if err != nil {
		return err
	}

	meta4.Published = &c.Args.Time.Time

	return c.Meta4.Put(meta4)
}
