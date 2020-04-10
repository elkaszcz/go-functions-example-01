// The cmd command starts an HTTP server.
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/GoogleCloudPlatform/golang-samples/functions/codelabs/gopher"
)

func main() {
	http.HandleFunc("/", gopher.Gopher)
	fmt.Println("Listening on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
