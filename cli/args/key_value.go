package args

import (
	"fmt"
	"strings"
)

type KeyValue struct {
	Key   string
	Value string
}

func MustNewKeyValue(value string) KeyValue {
	var a KeyValue

	err := a.UnmarshalFlag(value)
	if err != nil {
		panic(err)
	}

	return a
}

func (a *KeyValue) UnmarshalFlag(value string) error {
	parts := strings.SplitN(value, "=", 2)

	if len(parts) != 2 {
		return fmt.Errorf("Expected key=value format: %s", value)
	}

	a.Key = parts[0]
	a.Value = parts[1]

	return nil
}

func (a KeyValue) String() string {
	return fmt.Sprintf("%s=%s", a.Key, a.Value)
}
