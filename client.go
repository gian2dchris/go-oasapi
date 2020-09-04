package oasapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

/*
* (OK) url: "/api/?act=webGetLinesWithMLInfo",
* (OK) url: '/api/?act=webGetLangs',
* (OK) url: "/api/?act=webGetLines",
* (OK) url: '/api/?act=getLineName&p1='+lineCode,
* (OK) url: '/api/?act=webGetRoutes&p1='+lineCode,
* (OK) url: '/api/?act=getRoutesForLine&p1='+linecode,
* (OK) url: '/api/?act=webRouteDetails&p1='+routeCode,
* (OK) url: '/api/?act=getRouteName&p1='+routeCode,
* (OK) url: '/api/?act=webRoutesForStop&p1='+stopCode,
* (OK) url: '/api/?act=webGetRoutesDetailsAndStops&p1='+routeCode, JSON
* (OK) url: '/api/?act=webGetStops&p1='+routeCode,
* (OK) url: '/api/?act=getStopArrivals&p1='+stopCode,
* (OK) url: '/api/?act=getStopNameAndXY&p1='+stopCode,
* (OK) url: '/api/?act=getClosestStops&p1='+x+'&p2='+y,
* (OK) url: '/api/?act=getBusLocation&p1='+routeCode,
* (OK) url: '/api/?act=getScheduleDaysMasterline&p1='+lineCode,
* (OK) url: '/api/?act=getMLName&p1='+mlCode,
* url: '/api/?act=getDailySchedule&line_code='+lineCode, JSON

----------Not implemented----------
* Returns Null Responses ???
* url: "/api/?act=getPlacesTypes",
* url: '/api/?act=getPlaces&p1='+catCode,
* url: '/api/?act=getLinesAndRoutesForMl&p1='+mlCode,
* url: '/api/?act=getLinesAndRoutesForMlandLCode&p1='+mlCode+'&p2='+lid,
* url: '/api/?act=getSchedLines&p1='+mlCode+'&p2='+sdcCode+'&p3='+lineCode,
* url: '/api/?act=getEsavlDirections&p1=4&p2=3&p3=300&p4='+from[1]+'&p5='+from[0]+'&p6='+to[1]+'&p7='+to[0],
*/

const (
	baseURL = "http://telematics.oasa.gr/api/"
)

type Client struct {
	baseURL string
	// apiKey string
	HTTPClient *http.Client
}

func NewClient() *Client {

	return &Client{
		baseURL: baseURL,
		// apiKey: apiKey,
		HTTPClient: &http.Client{
			Timeout: time.Minute,
		},
	}
}

func (c *Client) sendRequest(req *http.Request, data interface{}) error {

	// API does not seem to take this into account, but anyway.
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Status Code %d", resp.StatusCode)
	}

	respBody, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return err
	}
	json.Unmarshal(respBody, data)
	// fmt.Println(string(respBody))
	return nil
}

type errorResponse struct {
	Error string `json:"error"`
}
