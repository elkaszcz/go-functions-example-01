package samplefunctions

import (
	"fmt"
	"net/http"
)

// HelloWorld prints "Hello, world."
func HelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world.")
}

// AlwaysError always gives an error
func AlwaysError(w http.ResponseWriter, r *http.Request) {
	http.Error(w, fmt.Sprintf("There was an error"), http.StatusInternalServerError)
	return
}
