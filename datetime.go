package uddf

import (
	"encoding/xml"
	"fmt"
	"time"
)

var (
	// DatetimeFormats is the list of all supported formats
	DatetimeFormats = []string{
		time.RFC3339,
		"2006-01",
		"200601",
		"2006-01-02",
		"20060102",
		"2006-01-02T15:04",
		"2006-01-02T15:04Z",
		"2006-01-02T15:04-07:00",
		"2006-01-02 15:04",
		"20060102T1504",
		"20060102T1504-0700",
		"T15:04",
		"T1504",
	}
)

// Datetime https://www.streit.cc/extern/uddf_v321/de/datetime.html
type Datetime time.Time

func (dt Datetime) String() string {
	return time.Time(dt).Format(time.RFC3339)
}

// MarshalXML is the custom XML marshaler for Datetime
func (dt Datetime) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(dt.String(), start)
}

// UnmarshalXML is the custom XML unmarshaler for Datetime
func (dt *Datetime) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var s string

	if err := d.DecodeElement(&s, &start); err != nil {
		return err
	}

	for _, f := range DatetimeFormats {
		t, err := time.Parse(f, s)
		if err != nil {
			continue
		}
		*dt = Datetime(t)
		return nil
	}

	return fmt.Errorf("%q is not a supported datetime format", s)
}
