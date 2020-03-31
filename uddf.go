// Package uddf implements the Universal Dive Data Format as in https://www.streit.cc/extern/uddf_v321/en/index.html
// use github.com/go-playground/validator for struct validation
package uddf

import (
	"encoding/xml"
)

// VERSION of the implemented UDDF spec
const VERSION = "3.2.1"

// New returns an UDDF document with the current version
func New() *UDDF {
	return &UDDF{
		Version: VERSION,
	}
}

// UDDF is the UDDF file
type UDDF struct {
	XMLName   xml.Name  `xml:"uddf"`
	Version   string    `xml:"version,attr"`
	Generator Generator `xml:"generator" validate:"required"`
}
