package main

import (
	"encoding/xml"
	"fmt"
	"time"

	"github.com/dpb587/blob-receipt/metalink"
)

func main() {
	// https://en.wikipedia.org/wiki/Metalink#Example_Metalink_4.0_.meta4_file
	published, _ := time.Parse(time.RFC3339, "2009-05-15T12:23:23Z")
	e1 := metalink.Metalink{
		Published: &published,
		File: []metalink.File{
			{
				Name:     "example.ext",
				Size:     14471447,
				Identity: "Example",
				Version:  "1.0",
				Language: []string{"en"},
				// Extension>Description: "A description of the example file for download."
				Hash: []metalink.Hash{
					{
						Type: "sha-256",
						Hash: "3d6fece8033d146d8611eab4f032df738c8c1283620fd02a1f2bfec6e27d590d",
					},
				},
				URL: []metalink.URL{
					{
						Location: "de",
						Priority: 1,
						URL:      "ftp://ftp.example.com/example.ext",
					},
					{
						Location: "fr",
						Priority: 1,
						URL:      "http://example.com/example.ext",
					},
				},
				MetaURL: []metalink.MetaURL{
					{
						MediaType: "torrent",
						Priority:  2,
						URL:       "http://example.com/example.ext.torrent",
					},
				},
			},
		},
	}

	e1b, err := xml.MarshalIndent(e1, "", "  ")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(e1b))

	fmt.Println("\n\n\n\n")

	e2s := `<?xml version="1.0" encoding="utf-8"?>
	<metalink xmlns="urn:ietf:params:xml:ns:metalink">
	<origin dynamic="true">https://curl.haxx.se/metalink.cgi?curl=tar.gz</origin>
	<updated>2017-03-23T02:34:59Z</updated>
	<generator>curl Metalink Generator</generator>
	<published>2017-02-24T07:50:50Z</published>
	<file name="curl-7.53.1.tar.gz">
	<publisher>
	 <name>curl</name>
	 <url>https://curl.haxx.se/</url>
	</publisher>
	<description>curl Generic source tar, gzip</description>
	<version>7.53.1</version>
	<size>3516153</size>
	<hash type="md5">9e49bb4cb275bf4464e7b69eb48613c0</hash>
	<hash type="sha-256">64f9b7ec82372edb8eaeded0a9cfa62334d8f98abc65487da01188259392911d</hash>
	<signature mediatype="application/pgp-signature">
	-----BEGIN PGP SIGNATURE-----

	iQEzBAABCgAdFiEEJ+3q8i86vOtQ25oSXMkI/bceEsIFAliv5cwACgkQXMkI/bce
	EsJ6FQgAiOlM+hfQut3C6FZALPRZDLa2XoMNuEmFGxhF0PrZELKfFRDsj5x6wbDg
	/jUUyR1yD2AIDWAIZj2Rv7b4jvGyUf+Gr3FyE4E56mETXZLyyLX8PSPGLrThCyB4
	IoHu4ufQEQX5Uc6k0LLuJOg4bp0f7huTdE6NKZC0ng4Ka3R5k/ISHE4fmmQpITv6
	/bXFO3dEdFyQEUVqUcwV5vape3IFzUnpqHBZkJ0UKe6wKEJlwFtvVA/8TGQLXLVV
	MhGYdZxjPv0LnShLcjyISPKraGGSSq3IdtT2YdU8ekuuWeHbX4v8Sobxh44DrKPw
	5nse+Ht5SdVjNFExvJn/4qfr9xwWyQ==
	=cCiX
	-----END PGP SIGNATURE-----
	</signature>
	<!-- resource preferences are for use in United states -->
	<url location="ca" priority="20">http://curl.mirror.anstey.ca/curl-7.53.1.tar.gz</url>
	<url location="de" priority="30">http://dl.uxnr.de/mirror/curl/curl-7.53.1.tar.gz</url>
	<url location="de" priority="30">https://dl.uxnr.de/mirror/curl/curl-7.53.1.tar.gz</url>
	<url location="se" priority="30">https://curl.haxx.se/download/curl-7.53.1.tar.gz</url>
	<url location="us" priority="10">http://curl.askapache.com/download/curl-7.53.1.tar.gz</url>
	</file>
	</metalink>`

	var e2 metalink.Metalink
	err = xml.Unmarshal([]byte(e2s), &e2)
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("%#+v", e2))

	e2rb, err := xml.MarshalIndent(e2, "", "  ")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(e2rb))

	fmt.Println("\n\n\n\n")

	e3s := ` <?xml version="1.0" encoding="UTF-8"?>
 <metalink xmlns="urn:ietf:params:xml:ns:metalink">
   <published>2009-05-15T12:23:23Z</published>
   <file name="example.ext">
     <size>14471447</size>
     <identity>Example</identity>
     <version>1.0</version>
     <language>en</language>
     <description>
     A description of the example file for download.
     </description>
     <hash type="sha-256">3d6fece8033d146d8611eab4f032df738c8c1283620fd02a1f2bfec6e27d590d</hash>
     <url location="de" priority="1">ftp://ftp.example.com/example.ext</url>
     <url location="fr" priority="1">http://example.com/example.ext</url>
     <metaurl mediatype="torrent" priority="2">http://example.com/example.ext.torrent</metaurl>
   </file>
 </metalink>
`

	var e3 metalink.Metalink
	err = xml.Unmarshal([]byte(e3s), &e3)
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("%#+v", e3))

	e3rb, err := xml.MarshalIndent(e3, "", "  ")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(e3rb))
}
