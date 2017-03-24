package cmd

import "github.com/dpb587/metalink"

type FileSetURL struct {
	Meta4File
	Location string         `long:"location" description:"ISO3166-1 country code for the geographical location"`
	Priority uint           `long:"priority" description:"Priority value between 1 and 999999. Lower values indicate a higher priority."`
	Args     FileSetURLArgs `positional-args:"true" required:"true"`
}

type FileSetURLArgs struct {
	URL string `positional-arg-name:"URL" description:"Download URI"`
}

func (c *FileSetURL) Execute(_ []string) error {
	file, err := c.Meta4File.Get()
	if err != nil {
		return err
	}

	for urlIdx, url := range file.URLs {
		if url.URL == c.Args.URL {
			file.URLs = append(file.URLs[:urlIdx], file.URLs[urlIdx+1:]...)

			break
		}
	}

	file.URLs = append(
		file.URLs,
		metalink.URL{
			Location: c.Location,
			Priority: c.Priority,
			URL:      c.Args.URL,
		},
	)

	return c.Meta4File.Put(file)
}
