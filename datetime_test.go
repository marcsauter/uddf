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
		"1985-02":                   "2006-01",
		"198502":                    "200601",
		"1997-01-05":                "2006-01-02",
		"19970105":                  "20060102",
		"2008-10-25T16:05":          "2006-01-02T15:04",
		"2008-10-25T16:05Z":         "2006-01-02T15:04Z",
		"2008-10-25T16:05+00:00":    "2006-01-02T15:04-07:00",
		"2008-10-25 16:05":          "2006-01-02 15:04",
		"20081025T1605":             "20060102T1504",
		"20081025T1605+0000":        "20060102T1504-0700",
		"T09:47":                    "T15:04",
		"T0947":                     "T1504",
		"2008-10-25T16:05:00+07:00": time.RFC3339,
		"2008-10-25T16:05:00Z":      time.RFC3339,
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

	t.Run("successful", func(t *testing.T) {
		exp := []byte(fmt.Sprintf(`<test><datetime>%s</datetime></test>`, now.Format(time.RFC3339)))
		act, err := xml.Marshal(&dt)
		assert.NoError(t, err)
		assert.Equal(t, exp, act)
	})
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

	t.Run("successful", func(t *testing.T) {
		for s, f := range DatetimeTests {
			ti, err := time.Parse(f, s)
			assert.NoError(t, err)
			exp.Datetime = uddf.Datetime(ti)
			act := &datetimeUnmarshalTest{}
			err = xml.Unmarshal([]byte(fmt.Sprintf(`<test><datetime>%s</datetime></test>`, s)), act)
			assert.NoError(t, err)
			assert.Equal(t, exp, act)
		}
	})

	t.Run("failure", func(t *testing.T) {
		act := &datetimeUnmarshalTest{}
		err := xml.Unmarshal([]byte(`<test><datetime>2020-03-29T3:21Z+02:00</datetime></test>`), act)
		assert.Error(t, err)
	})
}
