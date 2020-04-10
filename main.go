package main

import (
	"fmt"
	"log"
	"main/samplefunctions"
	"net/http"
)

func main() {
	http.HandleFunc("/HelloWorld", samplefunctions.HelloWorld)
	http.HandleFunc("/AlwaysError", samplefunctions.AlwaysError)
	fmt.Println("Listening on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
