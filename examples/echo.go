package main

import (
	"fmt"

	"github.com/hoanhan101/request"
	"github.com/hoanhan101/requesttest"
)

// Response describes a response from requesttest.Echo().
type Response struct {
	Method  string `json:"method"`
	Payload string `json:"payload"`
}

func main() {
	url, closer := requesttest.Echo()
	defer closer()

	r := new(Response)

	_ = request.GetJSON(
		&request.Options{
			URL:     url,
			Payload: map[string]string{"k1": "v1"},
		},
		r,
	)
	fmt.Printf("%+v\n", r)

	_ = request.PostJSON(
		&request.Options{
			URL:     url,
			Payload: map[string]string{"k1": "v1", "k2": "v2"},
		},
		r,
	)

	fmt.Printf("%+v\n", r)
}
