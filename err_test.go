package oasapi

import (
	"fmt"
	"net/http"
	"testing"
)

func TestErrorResponse(t *testing.T) {

	c := NewClient()

	req, err := http.NewRequest("GET", fmt.Sprintf("%s?act=notExists", baseURL), nil)
	if err != nil {
		fmt.Println(err)
	}

	var errorResponse errorResponse
	err = c.sendRequest(req, &errorResponse)
	if err != nil {
		t.Error("Error:", err)
	}
}
