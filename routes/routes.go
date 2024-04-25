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
	HandleArtistsRequests()
	fmt.Println("Server listening port", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))
}

func HandleArtistsRequests() {
	router.HandleFunc("/api/artists", artistController.GetAll).Methods(http.MethodGet)
	router.HandleFunc("/api/artists/{id}", artistController.GetById).Methods(http.MethodGet)
	router.HandleFunc("/api/artists", artistController.Create).Methods(http.MethodPost)
	router.HandleFunc("/api/artists/{id}", artistController.Update).Methods(http.MethodPut)
	router.HandleFunc("/api/artists/{id}", artistController.Delete).Methods(http.MethodDelete)
}
