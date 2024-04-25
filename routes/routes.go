package routes

import (
	"net/http"

	"github.com/gabrielsilper/golang-api-rest/controllers/artistController"
)

func HandleRequests() {
	http.HandleFunc("/api/artists", artistController.GetAll)
}
