package samplefunctions

import (
	"fmt"
	"net/http"
	"net/url"
)

// HelloWorld prints "Hello, world."
func HelloWorld(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		m, err := url.ParseQuery(r.URL.RawQuery)
		if err == nil {
			apiKey := m.Get("apiKey")
			fmt.Println(apiKey)
		}
		fmt.Fprint(w, "Hello World from GET!")
	} else {
		http.Error(w, "405 - Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

// AlwaysError always gives an error
func AlwaysError(w http.ResponseWriter, r *http.Request) {
	http.Error(w, fmt.Sprintf("There was an error"), http.StatusInternalServerError)
	return
}
