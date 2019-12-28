package requesttest_test

import (
	"fmt"

	"github.com/hoanhan101/request"
	"github.com/hoanhan101/requesttest"
)

// Response describes a response from requesttest.Echo().
type Response struct {
	Method  string      `json:"method"`
	Payload interface{} `json:"payload"`
}

func ExampleEcho() {
	url, closer := requesttest.Echo()
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
	// Output: &{Method:GET Payload:k1=v1}
}

func ExampleMock() {
	url, closer := requesttest.Mock(`{"method":"GET","payload":{"k1":"k2"}}`)
	defer closer()

	r := new(Response)

	err := request.GetJSON(
		&request.Options{
			URL: url,
		},
		r,
	)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%+v\n", r)
	// Output: &{Method:GET Payload:map[k1:k2]}
}
