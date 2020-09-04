package oasapi

import (
	"testing"
)

func TestWebGetLangs(t *testing.T) {

	c := NewClient()
	resp, err := c.getLangs()

	if err != nil {
		t.Error("Error:", err)
	}
	if resp == nil {
		t.Error("Error: null response")
	}
}
