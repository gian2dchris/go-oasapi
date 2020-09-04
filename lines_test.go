package oasapi

import (
	"fmt"
	"testing"
)

func TestWebGetLinesWithMLInfo(t *testing.T) {

	c := NewClient()
	resp, err := c.getLinesWithMLInfo()

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("LineID: %v\n", resp[13])
	}
}

func TestGetMLName(t *testing.T) {
	c := NewClient()
	resp, err := c.getMLName(73)

	if err != nil {
		fmt.Println(err)
	}

	if resp != nil {
		fmt.Println("MlDescr:", resp[0].MlDescr)
	} else {
		fmt.Println("Null Response:", resp)
	}
}

func TestWebGetLines(t *testing.T) {

	c := NewClient()
	resp, err := c.getLines()

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("LineCode: %+v\n", resp[0].LineCode)
	}
}

func TestWebGetLineName(t *testing.T) {

	c := NewClient()
	resp, err := c.getLineName(1151)

	if err != nil {
		fmt.Println(err)
	}

	if resp != nil {
		fmt.Printf("LineID: %s\n", resp[0].LineID)
	} else {
		fmt.Println("Null response:", resp)
	}
}

func TestGetScheduleDaysMasterline(t *testing.T) {
	c := NewClient()
	resp, err := c.getScheduleDaysMasterline(1151)

	if err != nil {
		fmt.Println(err)
	}

	if resp != nil {
		fmt.Println("SdcCode:", resp[0].SdcCode)
	} else {
		fmt.Println("Null Response:", resp)
	}
}
