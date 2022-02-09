package main

import (
	"fmt"
	"log"
	"net/http"
)

var port = "8080"

func main() {
	mux := routes()

	log.Println("Starting web server on port ", port)

	_ = http.ListenAndServe(fmt.Sprintf(":%s", port), mux)
}
