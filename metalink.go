package metalink

import (
	"encoding/xml"
	"time"
)

type Metalink struct {
	XMLName   xml.Name   `xml:"urn:ietf:params:xml:ns:metalink metalink" json:"-"`
	Files     []File     `xml:"file" json:"file,omitempty"`
	Generator string     `xml:"generator,,omitempty" json:"generator,omitempty"`
	Origin    *Origin    `xml:"origin,,omitempty" json:"origin,omitempty"`
	Published *time.Time `xml:"published,,omitempty" json:"published,omitempty"`
	Updated   *time.Time `xml:"updated,,omitempty" json:"updated,omitempty"`
}

type Origin struct {
	XMLName xml.Name `xml:"origin" json:"-"`
	Dynamic *bool    `xml:"dynamic,attr,omitempty" json:"dynamic,omitempty"`
	URL     string   `xml:",chardata" json:"url"`
}

type File struct {
	XMLName     xml.Name   `xml:"file" json:"-"`
	Name        string     `xml:"name,attr" json:"name"`
	Copyright   string     `xml:"copyright,,omitempty" json:"copyright,omitempty"`
	Description string     `xml:"description,,omitempty" json:"description,omitempty"`
	Hashes      []Hash     `xml:"hash,,omitempty" json:"hashes,omitempty"`
	Identity    string     `xml:"identity,,omitempty" json:"identity,omitempty"`
	Language    []string   `xml:"language,,omitempty" json:"language,omitempty"`
	Logo        string     `xml:"logo,,omitempty" json:"logo,omitempty"`
	MetaURLs    []MetaURL  `xml:"metaurl,,omitempty" json:"metaurl,omitempty"`
	OS          []string   `xml:"os,,omitempty" json:"os,omitempty"`
	Pieces      []Piece    `xml:"pieces,,omitempty" json:"piece,omitempty"`
	Publisher   *Publisher `xml:"publisher" json:"publisher,omitempty"`
	Signature   *Signature `xml:"signature" json:"signature,omitempty"`
	Size        uint64     `xml:"size,,omitempty" json:"size,omitempty"`
	URLs        []URL      `xml:"url,,omitempty" json:"url,omitempty"`
	Version     string     `xml:"version,omitempty" json:"version,omitempty"`
}

type URL struct {
	XMLName  xml.Name `xml:"url" json:"-"`
	Location string   `xml:"location,attr,omitempty" json:"location,omitempty"`
	Priority uint     `xml:"priority,attr,omitempty" json:"priority,omitempty"`
	URL      string   `xml:",chardata" json:"url"`
}

type Signature struct {
	XMLName   xml.Name `xml:"signature" json:"-"`
	MediaType string   `xml:"mediatype,attr" json:"mediatype"`
	Signature string   `xml:",cdata" json:"signature"`
}

type Publisher struct {
	XMLName xml.Name `xml:"publisher" json:"-"`
	Name    string   `xml:"name,attr" json:"name"`
	URL     string   `xml:"url,attr,omitempty" json:"url,omitempty"`
}

type Hash struct {
	XMLName xml.Name `xml:"hash" json:"-"`
	Type    string   `xml:"type,attr" json:"type"`
	Hash    string   `xml:",chardata" json:"hash"`
}

type Piece struct {
	XMLName xml.Name `xml:"pieces" json:"-"`
	Type    string   `xml:"type,attr" json:"type"`
	Length  string   `xml:"length,attr" json:"length"`
	Hash    []string `xml:"hash,chardata" json:"hash"`
}

type MetaURL struct {
	XMLName   xml.Name `xml:"metaurl" json:"-"`
	Priority  int      `xml:"priority,attr,omitempty" json:"priority,omitempty"`
	MediaType string   `xml:"mediatype,attr" json:"mediatype"`
	Name      string   `xml:"name,attr,omitempty" json:"name,omitempty"`
	URL       string   `xml:",chardata" json:"url"`
}

type Extra_ struct {
	XMLName xml.Name `json:"-"`
	Data    string   `xml:",innerxml"`
}
