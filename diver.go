package uddf

import "encoding/xml"

// Diver https://www.streit.cc/extern/uddf_v321/en/diver.html
type Diver struct {
	XMLName xml.Name `xml:"diver"`
	Owner   Owner    `xml:"owner"`
	Buddy   Buddy    `xml:"buddy"`
}

// Owner https://www.streit.cc/extern/uddf_v321/en/owner.html
type Owner struct{}

// Buddy https://www.streit.cc/extern/uddf_v321/en/buddy.html
type Buddy struct{}
