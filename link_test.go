package uddf_test

import (
	"encoding/xml"
	"testing"

	"github.com/marcsauter/uddf"
	"github.com/stretchr/testify/assert"
)

func TestMarshalLink(t *testing.T) {
	type linkMarshalTest struct {
		XMLName xml.Name  `xml:"test"`
		Link    uddf.Link `xml:"link"`
	}

	l := &linkMarshalTest{
		Link: uddf.Link{
			Reference: "testReference",
		},
	}

	t.Run("successful", func(t *testing.T) {
		exp := []byte(`<test><link ref="testReference"></link></test>`)
		act, err := xml.Marshal(&l)
		assert.NoError(t, err)
		t.Log(string(exp), string(act))
		assert.Equal(t, exp, act)
	})
}

func TestUnmarshalLink(t *testing.T) {
	type linkUnmarshalTest struct {
		XMLName xml.Name  `xml:"test"`
		Link    uddf.Link `xml:"link"`
	}

	exp := &linkUnmarshalTest{
		XMLName: xml.Name{
			Space: "",
			Local: "test",
		},
		Link: uddf.Link{
			XMLName: xml.Name{
				Space: "",
				Local: "link",
			},
			Reference: "testReference",
		},
	}

	t.Run("successful", func(t *testing.T) {
		act := &linkUnmarshalTest{}
		err := xml.Unmarshal([]byte(`<test><link ref="testReference"></link></test>`), act)
		assert.NoError(t, err)
		t.Log(exp, act)
		assert.Equal(t, exp, act)
	})

	t.Run("successful - self closing", func(t *testing.T) {
		act := &linkUnmarshalTest{}
		err := xml.Unmarshal([]byte(`<test><link ref="testReference"/></test>`), act)
		assert.NoError(t, err)
		t.Log(exp, act)
		assert.Equal(t, exp, act)
	})
}
