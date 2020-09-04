package oasapi

import (
	"fmt"
	"net/http"
)

type webGetStopsResponse []Stop

type Stop struct {
	StopCode       string `json:"StopCode"`
	StopID         string `json:"StopID"`
	StopDescr      string `json:"StopDescr"`
	StopDescrEng   string `json:"StopDescrEng"`
	StopStreet     string `json:"StopStreet"`
	StopStreetEng  string `json:"StopStreetEng"`
	StopHeading    string `json:"StopHeading"`
	StopLat        string `json:"StopLat"`
	StopLng        string `json:"StopLng"`
	RouteStopOrder string `json:"RouteStopOrder"`
	StopType       string `json:"StopType"`
	StopAmea       string `json:"StopAmea"`
}

func (c *Client) webGetStops(routeCode int) (webGetStopsResponse, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s?act=webGetStops&p1=%d", baseURL, routeCode), nil)
	if err != nil {
		return nil, err
	}

	var fullResponse webGetStopsResponse

	err = c.sendRequest(req, &fullResponse)
	if err != nil {
		return nil, err
	}

	return fullResponse, nil

}

type getStopNameAndXYResponse []StopNameAndXY

type StopNameAndXY struct {
	StopDescr          string `json:"stop_descr"`
	StopDescrMatrixEng string `json:"stop_descr_matrix_eng"`
	StopLat            string `json:"stop_lat"`
	StopLng            string `json:"stop_lng"`
	StopHeading        string `json:"stop_heading"`
	StopID             string `json:"stop_id"`
	IsTerminal         string `json:"isTerminal"`
}

func (c *Client) getStopNameAndXY(stopCode int) (getStopNameAndXYResponse, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/?act=getStopNameAndXY&p1=%d", baseURL, stopCode), nil)
	if err != nil {
		return nil, err
	}

	var fullResponse getStopNameAndXYResponse
	err = c.sendRequest(req, &fullResponse)
	if err != nil {
		return nil, err
	}
	return fullResponse, nil
}

type getClosestStopResponse []ClosestStop

type ClosestStop struct {
	StopCode      string `json:"StopCode"`
	StopID        string `json:"StopID"`
	StopDescr     string `json:"StopDescr"`
	StopDescrEng  string `json:"StopDescrEng"`
	StopStreet    string `json:"StopStreet"`
	StopStreetEng string `json:"StopStreetEng"`
	StopHeading   string `json:"StopHeading"`
	StopLat       string `json:"StopLat"`
	StopLng       string `json:"StopLng"`
	Distance      string `json:"distance"`
}

func (c *Client) getClosestStops(x float64, y float64) (getClosestStopResponse, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/?act=getClosestStops&p1=%.02f&p2=%.02f", baseURL, x, y), nil)
	if err != nil {
		return nil, err
	}

	var fullResponse getClosestStopResponse
	err = c.sendRequest(req, &fullResponse)
	if err != nil {
		return nil, err
	}

	return fullResponse, nil
}

type getStopArrivalsResponse []Arrival

type Arrival struct {
	RouteCode string `json:"route_code"`
	VehCode   string `json:"veh_code"`
	Btime2    string `json:"btime2"`
}

func (c *Client) getStopArrivals(stopCode int) (getStopArrivalsResponse, error) {

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/?act=getStopArrivals&p1=%d", baseURL, stopCode), nil)
	if err != nil {
		return nil, err
	}

	var fullResponse getStopArrivalsResponse
	err = c.sendRequest(req, &fullResponse)
	if err != nil {
		return nil, err
	}

	return fullResponse, nil
}
