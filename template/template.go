package template

import (
	"bytes"
	"text/template"

	"github.com/dpb587/metalink"
)

type Template struct {
	tmpl *template.Template
}

func New(uri string) (*Template, error) {
	tmpl := template.New("uri")

	tmpl, err := tmpl.Parse(uri)
	if err != nil {
		return nil, err
	}

	return &Template{tmpl: tmpl}, nil
}

func (t Template) ExecuteString(file metalink.File) (string, error) {
	wr := bytes.NewBuffer(nil)

	err := t.tmpl.Execute(wr, templateFile(file))
	if err != nil {
		return "", err
	}

	return wr.String(), nil
}
