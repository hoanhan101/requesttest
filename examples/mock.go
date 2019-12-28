package main

import (
	"fmt"

	"github.com/hoanhan101/request"
	"github.com/hoanhan101/requesttest"
)

// Response describes a sample response.
type Response struct {
	Status  string            `json:"status"`
	Payload map[string]string `json:"payload"`
}

func main() {
	url, closer := requesttest.Mock(`{"status":"ok","payload":{"k1":"v1"}}`)
	defer closer()

	r := new(Response)
	err := request.GetJSON(
		&request.Options{
			URL:     url,
			Payload: map[string]string{"k1": "v1"},
		},
		r,
	)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%+v\n", r)
}
