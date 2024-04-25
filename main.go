package main

import (
	"github.com/gabrielsilper/golang-api-rest/database"
	"github.com/gabrielsilper/golang-api-rest/routes"
)

const (
	port = 8080
)

func main() {
	database.ConnectDB()
	routes.StartServer(port)
}
