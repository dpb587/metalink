package args

import "strings"

type Filter struct {
	Type  string
	Value string
}

func (a *Filter) UnmarshalFlag(value string) error {
	split := strings.SplitN(value, ":", 2)

	a.Type = split[0]

	if len(split) > 1 {
		a.Value = split[1]
	}

	return nil
}
