package uddf

// Address https://www.streit.cc/extern/uddf_v321/en/address.html
type Address struct {
	Street   string `xml:"street,omitempty"`
	City     string `xml:"city,omitempty"`
	Postcode string `xml:"postcode,omitempty"`
	Province string `xml:"province,omitempty"`
	Country  string `xml:"country" validate:"required"`
}
