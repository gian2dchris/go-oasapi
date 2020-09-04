package oasapi

import (
	"fmt"
	"testing"
)

func TestWebGetRoutes(t *testing.T) {

	c := NewClient()
	resp, err := c.webGetRoutes(1151)

	if err != nil {
		fmt.Println(err)
	}

	if resp != nil {
		fmt.Printf("RouteCode: %s\n", resp[0].RouteCode)
	} else {
		fmt.Println("Null response:", resp)
	}
}

func TestGetRoutesForLine(t *testing.T) {

	c := NewClient()
	resp, err := c.getRoutesForLine(1151)

	if err != nil {
		fmt.Println(err)
	}

	if resp != nil {
		fmt.Printf("RouteID: %s\n", resp[0].RouteID)
	} else {
		fmt.Println("Null response:", resp)
	}
}
func TestWebRouteDetails(t *testing.T) {

	c := NewClient()
	resp, err := c.webRouteDetails(2484)

	if err != nil {
		fmt.Println(err)
	}

	if resp != nil {
		fmt.Printf("RoutedOrder: %s\n", resp[0].RoutedOrder)
	} else {
		fmt.Println("Null response:", resp)
	}

}
func TestGetRouteName(t *testing.T) {

	c := NewClient()
	resp, err := c.getRouteName(2484)

	if err != nil {
		fmt.Println(err)
	}
	if resp != nil {
		fmt.Printf("RouteDescr: %s\n", resp[0].RouteDescr)
	} else {
		fmt.Println("Null response:", resp)
	}
}

func TestWebGetStops(t *testing.T) {

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

func TestGetBusLocation(t *testing.T) {
	c := NewClient()
	resp, err := c.getBusLocation(2484)

	if err != nil {
		fmt.Println(err)
	}
	if resp != nil {
		fmt.Printf("VEH_NO: %s\n", resp[0].VEHNO)
	} else {
		fmt.Println("Null response:", resp)
	}

}

func TestRoutesForStop(t *testing.T) {
	c := NewClient()
	resp, err := c.webRoutesForStop(10175)

	if err != nil {
		fmt.Println(err)
	}
	if resp != nil {
		fmt.Printf("VEH_NO: %s\n", resp[0].RouteCode)
	} else {
		fmt.Println("Null response:", resp)
	}

}

func TestWebGetRoutesDetailsAndStops(t *testing.T) {
	c := NewClient()
	resp, err := c.webGetRouteDetailsAndStops(2484)

	if err != nil {
		fmt.Println(err)
	}

	if resp != nil {
		fmt.Println("RouteStops:", resp.Stops[0])
		fmt.Println("RouteDetails:", resp.Details[0])
	} else {
		fmt.Println("Null Response:", resp)
	}
}
