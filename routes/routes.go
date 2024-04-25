package routes

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gabrielsilper/golang-api-rest/controllers/artistController"
	"github.com/gorilla/mux"
)

var router = mux.NewRouter()

func StartServer(port int) {
	HandleRequests()
	fmt.Println("Server listening port", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))
}

func HandleRequests() {
	router.HandleFunc("/api/artists", artistController.GetAll).Methods(http.MethodGet)
	router.HandleFunc("/api/artists/{id}", artistController.GetById).Methods(http.MethodGet)
}
