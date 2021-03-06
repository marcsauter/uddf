package uddf

import "encoding/xml"

// Manufacturer https://www.streit.cc/extern/uddf_v321/en/manufacturer.html
type Manufacturer struct {
	XMLName   xml.Name `xml:"manufacturer"`
	ID        string   `xml:"id,attr" validate:"required"`
	Name      string   `xml:"name" validate:"required"`
	Aliasname string   `xml:"aliasname,omitempty"`
	Address   Address  `xml:"address,omitempty"`
	Contact   Contact  `xml:"contact,omitempty"`
}
