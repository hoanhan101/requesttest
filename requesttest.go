package requesttest

import (
	"fmt"
	"net/http"
	"net/http/httptest"
)

// Write a JSON response.
func Write(path string, response interface{}) (string, func()) {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)

	mux.HandleFunc(
		path,
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(200)
			fmt.Fprint(w, response)
		},
	)

	return fmt.Sprintf("%s%s", server.URL, path), server.Close
}

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
