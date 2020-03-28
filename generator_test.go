package uddf_test

import (
	"encoding/xml"
	"testing"

	"github.com/marcsauter/uddf"
	"github.com/stretchr/testify/assert"
)

func TestMarshalGeneratorType(t *testing.T) {
	type generatorTypeMarshalTest struct {
		XMLName xml.Name           `xml:"test"`
		Type    uddf.GeneratorType `xml:"type"`
	}

	gt := &generatorTypeMarshalTest{}

	t.Run("converter", func(t *testing.T) {
		gt.Type = "converter"
		exp := []byte(`<test><type>converter</type></test>`)
		act, err := xml.Marshal(&gt)
		assert.NoError(t, err)
		assert.Equal(t, exp, act)
	})

	t.Run("divecomputer", func(t *testing.T) {
		gt.Type = "divecomputer"
		exp := []byte(`<test><type>divecomputer</type></test>`)
		act, err := xml.Marshal(&gt)
		assert.NoError(t, err)
		assert.Equal(t, exp, act)
	})

	t.Run("logbook", func(t *testing.T) {
		gt.Type = "logbook"
		exp := []byte(`<test><type>logbook</type></test>`)
		act, err := xml.Marshal(&gt)
		assert.NoError(t, err)
		assert.Equal(t, exp, act)
	})

	t.Run("unknown", func(t *testing.T) {
		gt.Type = "unknown"
		act, err := xml.Marshal(&gt)
		assert.Error(t, err)
		assert.Nil(t, act)
	})
}

func TestUnmarshalGeneratorType(t *testing.T) {
	type generatorTypeUnmarshalTest struct {
		XMLName xml.Name           `xml:"test"`
		Type    uddf.GeneratorType `xml:"type"`
	}

	exp := &generatorTypeUnmarshalTest{
		XMLName: xml.Name{
			Space: "",
			Local: "test",
		},
	}

	t.Run("unknown", func(t *testing.T) {
		act := &generatorTypeUnmarshalTest{}
		err := xml.Unmarshal([]byte(`<test><type>unknown</type></test>`), act)
		assert.Error(t, err)
		assert.Equal(t, exp, act)
	})

	t.Run("converter", func(t *testing.T) {
		exp.Type = uddf.Converter
		act := &generatorTypeUnmarshalTest{}
		err := xml.Unmarshal([]byte(`<test><type>converter</type></test>`), act)
		assert.NoError(t, err)
		assert.Equal(t, exp, act)
	})

	t.Run("divecomputer", func(t *testing.T) {
		exp.Type = uddf.DiveComputer
		act := &generatorTypeUnmarshalTest{}
		err := xml.Unmarshal([]byte(`<test><type>divecomputer</type></test>`), act)
		assert.NoError(t, err)
		assert.Equal(t, exp, act)
	})

	t.Run("logbook", func(t *testing.T) {
		exp.Type = uddf.LogBook
		act := &generatorTypeUnmarshalTest{}
		err := xml.Unmarshal([]byte(`<test><type>logbook</type></test>`), act)
		assert.NoError(t, err)
		assert.Equal(t, exp, act)
	})
}
