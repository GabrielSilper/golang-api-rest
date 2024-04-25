package artistController

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gabrielsilper/golang-api-rest/models"
	"github.com/gorilla/mux"
)

func GetAll(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.ArtistSlice)
}

func GetById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var artist models.Artist

	for _, v := range models.ArtistSlice {
		if strconv.Itoa(v.ID) == id {
			artist = v
			break
		}
	}

	json.NewEncoder(w).Encode(artist)
}
