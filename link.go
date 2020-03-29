package uddf

import "encoding/xml"

// Link https://www.streit.cc/extern/uddf_v321/en/link.html
type Link struct {
	XMLName   xml.Name `xml:"link"`
	Reference string   `xml:"ref,attr"`
}
