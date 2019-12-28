package requesttest_test

import (
	"fmt"

	"github.com/hoanhan101/request"
)

// Response describes a response from requesttest.Echo().
type Response struct {
	Method  string `json:"method"`
	Payload string `json:"payload"`
}

func ExampleEcho() {
	url, closer := Echo()
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
