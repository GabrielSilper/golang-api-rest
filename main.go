package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gabrielsilper/golang-api-rest/routes"
)

const (
	port = 8000
)

func startServer() {
	routes.HandleRequests()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}

func main() {
	fmt.Println("Server listening port", port)
	startServer()
}
