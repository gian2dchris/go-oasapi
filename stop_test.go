package oasapi

import (
	"fmt"
	"testing"
)

func TestWebGetStop(t *testing.T) {

	c := NewClient()
	resp, err := c.webGetStops(2484)

	if err != nil {
		fmt.Println(err)
	}

	if resp != nil {
		fmt.Printf("StopCode: %s\n", resp[0].StopCode)
	} else {
		fmt.Println("Null response:", resp)
	}

}

func TestGetStopNameAndXY(t *testing.T) {
	c := NewClient()
	resp, err := c.getStopNameAndXY(10175)

	if err != nil {
		fmt.Println(err)
	}

	if resp != nil {
		fmt.Println("StopID:", resp[0].StopID)
	} else {
		fmt.Println("Null Response:", resp)
	}

}

func TestGetClosestStops(t *testing.T) {
	c := NewClient()
	resp, err := c.getClosestStops(40, 25)

	if err != nil {
		fmt.Println(err)
	}

	if resp != nil {
		fmt.Println("StopCode:", resp[0].StopCode)
	} else {
		fmt.Println("Null Response:", resp)
	}

}
func TestGetStopArrivals(t *testing.T) {
	c := NewClient()
	resp, err := c.getStopArrivals(10175)

	if err != nil {
		fmt.Println(err)
	}
	if resp != nil {
		fmt.Printf("RouteCode: %s\n", resp[0].RouteCode)
	} else {
		fmt.Println("Null response:", resp)
	}
}
