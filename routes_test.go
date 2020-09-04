package oasapi

import (
	"testing"
)

func TestWebGetRoutes(t *testing.T) {

	c := NewClient()
	resp, err := c.webGetRoutes(1151)

	if err != nil {
		t.Error("Error:", err)
	}
	if resp == nil {
		t.Error("Error: null response")
	}
}

func TestGetRoutesForLine(t *testing.T) {

	c := NewClient()
	resp, err := c.getRoutesForLine(1151)

	if err != nil {
		t.Error("Error:", err)
	}
	if resp == nil {
		t.Error("Error: null response")
	}
}
func TestWebRouteDetails(t *testing.T) {

	c := NewClient()
	resp, err := c.webRouteDetails(2484)

	if err != nil {
		t.Error("Error:", err)
	}
	if resp == nil {
		t.Error("Error: null response")
	}

}
func TestGetRouteName(t *testing.T) {

	c := NewClient()
	resp, err := c.getRouteName(2484)

	if err != nil {
		t.Error("Error:", err)
	}
	if resp == nil {
		t.Error("Error: null response")
	}
}

func TestWebGetStops(t *testing.T) {

	c := NewClient()
	resp, err := c.webGetStops(2484)

	if err != nil {
		t.Error("Error:", err)
	}
	if resp == nil {
		t.Error("Error: null response")
	}
}

func TestGetBusLocation(t *testing.T) {
	c := NewClient()
	resp, err := c.getBusLocation(2484)

	if err != nil {
		t.Error("Error:", err)
	}
	if resp == nil {
		t.Error("Error: null response")
	}

}

func TestRoutesForStop(t *testing.T) {
	c := NewClient()
	resp, err := c.webRoutesForStop(10175)

	if err != nil {
		t.Error("Error:", err)
	}
	if resp == nil {
		t.Error("Error: null response")
	}

}

func TestWebGetRoutesDetailsAndStops(t *testing.T) {
	c := NewClient()
	resp, err := c.webGetRouteDetailsAndStops(2484)

	if err != nil {
		t.Error("Error:", err)
	}
	if resp == nil {
		t.Error("Error: null response")
	}
}
