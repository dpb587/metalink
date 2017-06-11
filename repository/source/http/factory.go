package http

import (
	gohttp "net/http"

	"github.com/dpb587/metalink/repository/source"
)

type Factory struct{}

var _ source.Factory = &Factory{}

func NewFactory() Factory {
	return Factory{}
}

func (f Factory) Schemes() []string {
	return []string{"http", "https"}
}

func (f Factory) Create(uri string, _ map[string]interface{}) (source.Source, error) {
	return NewSource(uri, &gohttp.Client{}), nil
}
