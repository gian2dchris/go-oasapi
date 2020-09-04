package oasapi

import (
	"fmt"
	"net/http"
)

type webGetLinesWithMLInfoResponse []LineWithMLInfo

type LineWithMLInfo struct {
	MlCode       string `json:"ml_code"`
	SdcCode      string `json:"sdc_code"`
	LineCode     string `json:"line_code"`
	LineID       string `json:"line_id"`
	LineDescr    string `json:"line_descr"`
	LineDescrEng string `json:"line_descr_eng"`
	MldMaster    string `json:"mld_master"`
}

func (c *Client) getLinesWithMLInfo() (webGetLinesWithMLInfoResponse, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s?act=webGetLinesWithMLInfo", baseURL), nil)
	if err != nil {
		return nil, err
	}

	var fullResponse webGetLinesWithMLInfoResponse

	err = c.sendRequest(req, &fullResponse)
	if err != nil {
		return nil, err
	}

	return fullResponse, nil
}

type getMLNameResponse []MlInfo

type MlInfo struct {
	MlDescr    string      `json:"ml_descr"`
	MlDescrEng interface{} `json:"ml_descr_eng"`
}

func (c *Client) getMLName(mlCode int) (getMLNameResponse, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s?act=getMLName&p1=%d", baseURL, mlCode), nil)
	if err != nil {
		return nil, err
	}

	var fullResponse getMLNameResponse

	err = c.sendRequest(req, &fullResponse)
	if err != nil {
		return nil, err
	}
	return fullResponse, nil
}

type webGetLinesResponse []Line

type Line struct {
	LineCode     string `json:"LineCode"`
	LineID       string `json:"LineID"`
	LineDescr    string `json:"LineDescr"`
	LineDescrEng string `json:"LineDescrEng"`
}

func (c *Client) getLines() (webGetLinesResponse, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s?act=webGetLines", baseURL), nil)
	if err != nil {
		return nil, err
	}

	var fullResponse webGetLinesResponse

	err = c.sendRequest(req, &fullResponse)
	if err != nil {
		return nil, err
	}

	return fullResponse, nil
}

type webGetLineNameResponse []LineName

type LineName struct {
	LineDescr    string `json:"line_descr"`
	LineDescrEng string `json:"line_descr_eng"`
	LineID       string `json:"line_id"`
}

func (c *Client) getLineName(lineCode int) (webGetLineNameResponse, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s?act=getLineName&p1=%d", baseURL, lineCode), nil)
	if err != nil {
		return nil, err
	}

	var fullResponse webGetLineNameResponse

	err = c.sendRequest(req, &fullResponse)
	if err != nil {
		return nil, err
	}

	return fullResponse, nil
}

type getScheduleDaysMasterlineResponse []ScheduleDay

type ScheduleDay struct {
	SdcDescr    string `json:"sdc_descr"`
	SdcDescrEng string `json:"sdc_descr_eng"`
	SdcCode     string `json:"sdc_code"`
	string      `json:""`
}

func (c *Client) getScheduleDaysMasterline(lineCode int) (getScheduleDaysMasterlineResponse, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s?act=getScheduleDaysMasterline&p1=%d", baseURL, lineCode), nil)
	if err != nil {
		return nil, err
	}

	var fullResponse getScheduleDaysMasterlineResponse

	err = c.sendRequest(req, &fullResponse)
	if err != nil {
		return nil, err
	}
	return fullResponse, nil
}
