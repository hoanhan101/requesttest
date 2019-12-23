package requesttest

import (
	"fmt"
	"net/http"
	"net/http/httptest"
)

// Serve registers a route, overwrites its response and returns a ready-to-use URL.
//
// For example:
//  url, closer := requesttest.Serve("/get", `{"status":"ok"}`)
//  defer closer()
func Serve(path string, response interface{}) (string, func()) {
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
