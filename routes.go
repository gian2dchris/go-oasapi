package oasapi

import (
	"fmt"
	"net/http"
)

type webGetRoutesResponse []Route

type Route struct {
	RouteCode     string `json:"RouteCode"`
	LineCode      string `json:"LineCode"`
	RouteDescr    string `json:"RouteDescr"`
	RouteDescrEng string `json:"RouteDescrEng"`
	RouteType     string `json:"RouteType"`
	RouteDistance string `json:"RouteDistance"`
}

func (c *Client) webGetRoutes(lineCode int) (webGetRoutesResponse, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s?act=webGetRoutes&p1=%d", baseURL, lineCode), nil)
	if err != nil {
		return nil, err
	}

	var fullResponse webGetRoutesResponse

	err = c.sendRequest(req, &fullResponse)
	if err != nil {
		return nil, err
	}

	return fullResponse, nil

}

type getRoutesForLineResponse []RouteForLine

type RouteForLine struct {
	RouteCode     string `json:"route_code"`
	RouteID       string `json:"route_id"`
	RouteDescr    string `json:"route_descr"`
	RouteActive   string `json:"route_active"`
	RouteDescrEng string `json:"route_descr_eng"`
}

func (c *Client) getRoutesForLine(lineCode int) (getRoutesForLineResponse, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s?act=getRoutesForLine&p1=%d", baseURL, lineCode), nil)
	if err != nil {
		return nil, err
	}

	var fullResponse getRoutesForLineResponse

	err = c.sendRequest(req, &fullResponse)
	if err != nil {
		return nil, err
	}

	return fullResponse, nil

}

type webRouteDetailsResponse []RouteDetails

type RouteDetails struct {
	RoutedX     string `json:"routed_x"`
	RoutedY     string `json:"routed_y"`
	RoutedOrder string `json:"routed_order"`
}

func (c *Client) webRouteDetails(routeCode int) (webRouteDetailsResponse, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s?act=webRouteDetails&p1=%d", baseURL, routeCode), nil)
	if err != nil {
		return nil, err
	}

	var fullResponse webRouteDetailsResponse

	err = c.sendRequest(req, &fullResponse)
	if err != nil {
		return nil, err
	}

	return fullResponse, nil

}

type getRouteNameResponse []RouteName

type RouteName struct {
	RouteDescr        string `json:"route_descr"`
	RouteDepartureEng string `json:"route_departure_eng"`
}

func (c *Client) getRouteName(routeCode int) (getRouteNameResponse, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s?act=getRouteName&p1=%d", baseURL, routeCode), nil)
	if err != nil {
		return nil, err
	}

	var fullResponse getRouteNameResponse

	err = c.sendRequest(req, &fullResponse)
	if err != nil {
		return nil, err
	}

	return fullResponse, nil

}

type getBusLocationResponse []BusLocation

type BusLocation struct {
	VEHNO     string `json:"VEH_NO"`
	CSDATE    string `json:"CS_DATE"`
	CSLAT     string `json:"CS_LAT"`
	CSLNG     string `json:"CS_LNG"`
	ROUTECODE string `json:"ROUTE_CODE"`
}

func (c *Client) getBusLocation(routeCode int) (getBusLocationResponse, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s?act=getBusLocation&p1=%d", baseURL, routeCode), nil)
	if err != nil {
		return nil, err
	}

	var fullResponse getBusLocationResponse

	err = c.sendRequest(req, &fullResponse)
	if err != nil {
		return nil, err
	}

	return fullResponse, nil

}

type webRoutesForStop []RouteAndLine

type RouteAndLine struct {
	RouteCode      string `json:"RouteCode"`
	LineCode       string `json:"LineCode"`
	Hidden         string `json:"hidden"`
	RouteDescr     string `json:"RouteDescr"`
	RouteDescrEng  string `json:"RouteDescrEng"`
	RouteType      string `json:"RouteType"`
	RouteDistance  string `json:"RouteDistance"`
	LineID         string `json:"LineID"`
	LineDescr      string `json:"LineDescr"`
	LineDescrEng   string `json:"LineDescrEng"`
	MasterLineCode string `json:"MasterLineCode"`
}

func (c *Client) webRoutesForStop(stopCode int) (webRoutesForStop, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s?act=webRoutesForStop&p1=%d", baseURL, stopCode), nil)
	if err != nil {
		return nil, err
	}

	var fullResponse webRoutesForStop

	err = c.sendRequest(req, &fullResponse)
	if err != nil {
		return nil, err
	}

	return fullResponse, nil

}

type webGetRouteDetailsAndStopsResponse struct {
	Details []StopDetails `json:"details"`
	Stops   []Stop        `json:"stops"`
}

type StopDetails struct {
	RoutedX     string `json:"routed_x"`
	RoutedY     string `json:"routed_y"`
	RoutedOrder string `json:"routed_order"`
}

func (c *Client) webGetRouteDetailsAndStops(routeCode int) (*webGetRouteDetailsAndStopsResponse, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s?act=webGetRoutesDetailsAndStops&p1=%d", baseURL, routeCode), nil)
	if err != nil {
		return nil, err
	}

	var fullResponse webGetRouteDetailsAndStopsResponse

	err = c.sendRequest(req, &fullResponse)
	if err != nil {
		return nil, err
	}

	return &fullResponse, nil
}
