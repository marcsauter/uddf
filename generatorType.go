package uddf

import (
	"encoding/xml"
	"fmt"
)

// Valid generator types
const (
	Converter    GeneratorType = "converter"
	DiveComputer GeneratorType = "divecomputer"
	LogBook      GeneratorType = "logbook"
)

// GeneratorType https://www.streit.cc/extern/uddf_v321/en/generator.html
type GeneratorType string

func (gt GeneratorType) String() string {
	return string(gt)
}

// MarshalXML is the custom XML marshaler for GeneratorType
func (gt GeneratorType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	switch gt {
	case Converter, DiveComputer, LogBook:
		return e.EncodeElement(string(gt), start)
	default:
		return fmt.Errorf("generator type has to be one of converter, divecomputer, logbook")
	}
}

// UnmarshalXML is the custom XML unmarshaler for GeneratorType
func (gt *GeneratorType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var s string

	if err := d.DecodeElement(&s, &start); err != nil {
		return err
	}

	switch s {
	case "converter", "divecomputer", "logbook":
		*gt = GeneratorType(s)
		return nil
	default:
		return fmt.Errorf("generator type has to be one of converter, divecomputer, logbook")
	}
}
