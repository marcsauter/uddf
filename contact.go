package uddf

// Contact https://www.streit.cc/extern/uddf_v321/en/contact.html
type Contact struct {
	Email       string `xml:"email,omitempty" validate:"email"`
	Fax         string `xml:"fax,omitempty"`
	Homepage    string `xml:"homepage,omitempty" validate:"uri"`
	Language    string `xml:"language,omitempty"`
	Mobilephone string `xml:"mobilephone,omitempty"`
	Phone       string `xml:"phone,omitempty"`
}
