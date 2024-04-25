package artistController

import (
	"encoding/json"
	"net/http"

	artistService "github.com/gabrielsilper/golang-api-rest/services/artistService"
	"github.com/gorilla/mux"
)

func GetAll(w http.ResponseWriter, r *http.Request) {
	artists := artistService.FindAll()
	json.NewEncoder(w).Encode(artists)
}

func GetById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	artist := artistService.FindById(id)

	json.NewEncoder(w).Encode(artist)
}
