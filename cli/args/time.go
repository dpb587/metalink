package args

import (
	"time"

	"github.com/pkg/errors"
)

type Time struct {
	time.Time
}

func MustNewTime(value string) Time {
	var a Time

	err := a.UnmarshalFlag(value)
	if err != nil {
		panic(err)
	}

	return a
}

func (a *Time) UnmarshalFlag(value string) error {
	parsed, err := time.Parse(time.RFC3339, value)
	if err != nil {
		return errors.Wrap(err, "Parsing time argument")
	}

	a.Time = parsed

	return nil
}

func (a Time) String() string {
	return a.Time.Format(time.RFC3339)
}
