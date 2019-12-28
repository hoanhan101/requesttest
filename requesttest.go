package requesttest

import (
	"fmt"
	"net/http"
	"net/http/httptest"
)

// Echo a request's method and payload in JSON.
func Echo() (string, func()) {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)

	mux.HandleFunc(
		"/",
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(200)

			_ = r.ParseForm()
			fmt.Fprint(
				w,
				fmt.Sprintf(
					`{"method":"%s","payload":"%s"}`,
					r.Method,
					r.Form.Encode(),
				),
			)
		},
	)

	return server.URL, server.Close
}

// Mock a JSON response.
func Mock(response interface{}) (string, func()) {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)

	mux.HandleFunc(
		"/",
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(200)
			fmt.Fprint(w, response)
		},
	)

	return server.URL, server.Close
}
