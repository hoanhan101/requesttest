package main

import (
	"fmt"

	"github.com/hoanhan101/request"
	"github.com/hoanhan101/requesttest"
)

// Response is a sample response.
type Response struct {
	Status string            `json:"status"`
	Query  map[string]string `json:"query"`
}

func main() {
	url, closer := requesttest.Serve("/get", `{"status":"ok","query":{"k1":"v1"}}`)
	defer closer()

	r := new(Response)
	err := request.GetJSON(url, nil, r)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%+v\n", r)
}
