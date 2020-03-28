package uddf_test

import (
	"encoding/xml"
	"fmt"
	"testing"
	"time"

	"github.com/marcsauter/uddf"
	"github.com/stretchr/testify/assert"
)

var (
	DatetimeTests = map[string]string{
		"2006-01":                "1985-02",
		"200601":                 "198502",
		"2006-01-02":             "1997-01-05",
		"20060102":               "19970105",
		"2006-01-02T15:04":       "2008-10-25T16:05",
		"2006-01-02T15:04Z":      "2008-10-25T16:05Z",
		"2006-01-02T15:04-07:00": "2008-10-25T16:05+00:00",
		"2006-01-02 15:04":       "2008-10-25 16:05",
		"20060102T1504":          "20081025T1605",
		"20060102T1504-0700":     "20081025T1605+0000",
		"T15:04":                 "T09:47",
		"T1504":                  "T0947",
		//		time.RFC3339:             "2008-10-25T16:05:00Z07:00",
	}
)

func TestMarshalDatetime(t *testing.T) {
	type datetimeMarshalTest struct {
		XMLName  xml.Name      `xml:"test"`
		Datetime uddf.Datetime `xml:"datetime"`
	}

	now := time.Now()
	dt := &datetimeMarshalTest{
		Datetime: uddf.Datetime(now),
	}
	exp := []byte(fmt.Sprintf(`<test><datetime>%s</datetime></test>`, now.Format(time.RFC3339)))
	act, err := xml.Marshal(&dt)
	assert.NoError(t, err)
	assert.Equal(t, exp, act)
}

func TestUnmarshalDatetime(t *testing.T) {
	type datetimeUnmarshalTest struct {
		XMLName  xml.Name      `xml:"test"`
		Datetime uddf.Datetime `xml:"datetime"`
	}

	exp := &datetimeUnmarshalTest{
		XMLName: xml.Name{
			Space: "",
			Local: "test",
		},
	}

	for f, s := range DatetimeTests {
		ti, err := time.Parse(f, s)
		assert.NoError(t, err)
		exp.Datetime = uddf.Datetime(ti)
		act := &datetimeUnmarshalTest{}
		err = xml.Unmarshal([]byte(fmt.Sprintf(`<test><datetime>%s</datetime></test>`, s)), act)
		assert.NoError(t, err)
		assert.Equal(t, exp, act)
	}
}
