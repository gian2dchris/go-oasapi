package oasapi

import (
	"testing"
)

func TestWebGetLinesWithMLInfo(t *testing.T) {

	c := NewClient()
	resp, err := c.getLinesWithMLInfo()

	if err != nil {
		t.Error("Error:", err)
	}
	if resp == nil {
		t.Error("Error: null response")
	}
}

func TestGetMLName(t *testing.T) {
	c := NewClient()
	resp, err := c.getMLName(73)

	if err != nil {
		t.Error("Error:", err)
	}
	if resp == nil {
		t.Error("Error: null response")
	}
}

func TestWebGetLines(t *testing.T) {

	c := NewClient()
	resp, err := c.getLines()

	if err != nil {
		t.Error("Error:", err)
	}
	if resp == nil {
		t.Error("Error: null response")
	}
}

func TestWebGetLineName(t *testing.T) {

	c := NewClient()
	resp, err := c.getLineName(1151)

	if err != nil {
		t.Error("Error:", err)
	}
	if resp == nil {
		t.Error("Error: null response")
	}
}

func TestGetScheduleDaysMasterline(t *testing.T) {
	c := NewClient()
	resp, err := c.getScheduleDaysMasterline(1151)

	if err != nil {
		t.Error("Error:", err)
	}
	if resp == nil {
		t.Error("Error: null response")
	}
}
