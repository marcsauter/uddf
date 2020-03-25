package uddf_test

import (
	"encoding/xml"
	"testing"

	"github.com/marcsauter/uddf"
	"github.com/stretchr/testify/assert"
)

func TestMarshal(t *testing.T) {
	d := uddf.New()
	b, err := xml.Marshal(d)
	assert.NoError(t, err)
	t.Log(string(b))
}

func TestUnmarshal(t *testing.T) {

}
