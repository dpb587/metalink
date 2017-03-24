package cmd

import "fmt"

type Files struct {
	Meta4
}

func (c *Files) Execute(_ []string) error {
	meta4, err := c.Meta4.Get()
	if err != nil {
		return err
	}

	for _, file := range meta4.Files {
		fmt.Println(file.Name)
	}

	return nil
}
