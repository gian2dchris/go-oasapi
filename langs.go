package oasapi

import (
	"fmt"
	"net/http"
)

type webGetLangsResponse []Lang

type Lang struct {
	LangID string `json:"lang_id"`
	El     string `json:"el"`
	En     string `json:"en"`
}

//webGetLangs: asd
func (c *Client) getLangs() (webGetLangsResponse, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s?act=webGetLangs", baseURL), nil)
	if err != nil {
		return nil, err
	}

	var fullResponse webGetLangsResponse

	err = c.sendRequest(req, &fullResponse)
	if err != nil {
		return nil, err
	}
	return fullResponse, nil
}
