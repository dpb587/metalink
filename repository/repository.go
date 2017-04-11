package repository

import (
	"encoding/xml"

	"github.com/dpb587/metalink"
)

type RepositoryMetalink struct {
	Reference RepositoryMetalinkReference `xml:"https://dpb587.github.io/metalink-repository/schema-0.1.0.xsd repository-reference,,omitempty" json:"repository_reference,omitempty"`
	metalink.Metalink
}

type RepositoryMetalinkReference struct {
	Repository string `xml:"repository,,omitempty" json:"repository"`
	Path       string `xml:"path,,omitempty" json:"path"`
	Version    string `xml:"version,,omitempty" json:"version"`
}

type Repository struct {
	XMLName   xml.Name             `xml:"https://dpb587.github.io/metalink-repository/schema-0.1.0.xsd repository" json:"-"`
	Metalinks []RepositoryMetalink `xml:"metalink" json:"metalink"`
}
