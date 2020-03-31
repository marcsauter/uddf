package uddf

import "encoding/xml"

// Generator https://www.streit.cc/extern/uddf_v321/en/generator.html
type Generator struct {
	XMLName      xml.Name      `xml:"generator"`
	Name         string        `xml:"name" validate:"required"`
	Aliasname    string        `xml:"aliasname,omitempty"`
	Manufacturer Manufacturer  `xml:"manufacturer,omitempty"`
	Version      string        `xml:"version,omitempty"`
	Datetime     Datetime      `xml:"datetime,omitempty"`
	Type         GeneratorType `xml:"type,omitempty"`
	Link         Link          `xml:"link,omitempty"`
}
