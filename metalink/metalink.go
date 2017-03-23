package metalink

import (
	"encoding/xml"
	"time"
)

type Metalink struct {
	XMLName xml.Name `xml:"urn:ietf:params:xml:ns:metalink metalink"`

	File      []File     `xml:"file"`
	Generator string     `xml:"generator,,omitempty"`
	Origin    *Origin    `xml:"origin,,omitempty"`
	Published *time.Time `xml:"published,,omitempty"`
	Updated   *time.Time `xml:"updated,,omitempty"`
}

type Origin struct {
	XMLName xml.Name `xml:"origin"`

	Dynamic *bool  `xml:"dynamic,attr,omitempty"`
	URL     string `xml:",chardata"`
}

type File struct {
	XMLName xml.Name `xml:"file"`

	Name        string     `xml:"name,attr"`
	Copyright   string     `xml:"copyright,,omitempty"`
	Description string     `xml:"description,,omitempty"`
	Hash        []Hash     `xml:"hash,,omitempty"`
	Identity    string     `xml:"identity,,omitempty"`
	Language    []string   `xml:"language,,omitempty"`
	Logo        string     `xml:"logo,,omitempty"`
	MetaURL     []MetaURL  `xml:"metaurl,,omitempty"`
	OS          []string   `xml:"os,,omitempty"`
	Pieces      []Piece    `xml:"pieces,,omitempty"`
	Publisher   *Publisher `xml:"publisher"`
	Signature   *Signature `xml:"signature"`
	Size        uint64     `xml:"size,,omitempty"`
	URL         []URL      `xml:"url,,omitempty"`
	Version     string     `xml:"version,omitempty"`
}

type URL struct {
	XMLName xml.Name `xml:"url"`

	Location string `xml:"location,attr,omitempty"`
	Priority uint   `xml:"priority,attr,omitempty"`
	URL      string `xml:",chardata"`
}
type Signature struct {
	XMLName xml.Name `xml:"signature"`

	MediaType string `xml:"mediatype,attr"`
	Signature string `xml:",cdata"`
}

type Publisher struct {
	XMLName xml.Name `xml:"publisher"`

	Name string `xml:"name,attr"`
	URL  string `xml:"url,attr,omitempty"`
}

type Hash struct {
	XMLName xml.Name `xml:"hash"`

	Type string `xml:"type,attr"`
	Hash string `xml:",chardata"`
}

type Piece struct {
	XMLName xml.Name `xml:"pieces"`

	Type   string   `xml:"type,attr"`
	Length string   `xml:"length,attr"`
	Hash   []string `xml:"hash,chardata"`
}

type MetaURL struct {
	XMLName xml.Name `xml:"metaurl"`

	Priority  int    `xml:"priority,attr,omitempty"`
	MediaType string `xml:"mediatype,attr"`
	Name      string `xml:"name,attr,omitempty"`
	URL       string `xml:",chardata"`
}
