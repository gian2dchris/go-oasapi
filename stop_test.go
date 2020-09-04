package oasapi

import (
	"testing"
)

func TestWebGetStop(t *testing.T) {

	c := NewClient()
	resp, err := c.webGetStops(2484)

	if err != nil {
		t.Error("Error:", err)
	}
	if resp == nil {
		t.Error("Error: null response")
	}

}

func TestGetStopNameAndXY(t *testing.T) {
	c := NewClient()
	resp, err := c.getStopNameAndXY(10175)

	if err != nil {
		t.Error("Error:", err)
	}
	if resp == nil {
		t.Error("Error: null response")
	}

}

func TestGetClosestStops(t *testing.T) {
	c := NewClient()
	resp, err := c.getClosestStops(40, 25)

	if err != nil {
		t.Error("Error:", err)
	}
	if resp == nil {
		t.Error("Error: null response")
	}

}
func TestGetStopArrivals(t *testing.T) {
	c := NewClient()
	resp, err := c.getStopArrivals(10175)

	if err != nil {
		t.Error("Error:", err)
	}
	if resp == nil {
		t.Error("Error: null response")
	}
}
