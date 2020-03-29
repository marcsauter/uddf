package uddf

import "encoding/xml"

// Manufacturer https://www.streit.cc/extern/uddf_v321/en/manufacturer.html
type Manufacturer struct {
	XMLName   xml.Name `xml:"manufacturer"`
	Name      string   `xml:"name"`
	Aliasname string   `xml:"aliasname,omitempty"`
	Address   Address  `xml:"address,omitempty"`
	Contact   Contact  `xml:"contact,omitempty"`
}
