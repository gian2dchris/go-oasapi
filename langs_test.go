package oasapi

import (
	"fmt"
	"testing"
)

func TestWebGetLangs(t *testing.T) {

	c := NewClient()
	resp, err := c.getLangs()

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("LangID: %+v\n", resp[0].LangID)
	}
}
