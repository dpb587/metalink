package metalink

import (
	"encoding/xml"

	"github.com/pkg/errors"
)

func Unmarshal(data []byte, meta4 *Metalink) error {
	err := xml.Unmarshal(data, meta4)
	if err != nil {
		return errors.Wrap(err, "Unmarshaling XML")
	}

	return nil
}

func Marshal(r Metalink) ([]byte, error) {
	data, err := xml.MarshalIndent(r, "", "  ")
	if err != nil {
		return nil, errors.Wrap(err, "Marshaling XML")
	}

	return append(data, '\n'), nil
}
