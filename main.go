package main

import (
	"github.com/gabrielsilper/golang-api-rest/routes"
)

const (
	port = 8080
)

func main() {
	routes.StartServer(port)
}
