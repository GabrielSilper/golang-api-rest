package artistController

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gabrielsilper/golang-api-rest/models"
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

func Create(w http.ResponseWriter, r *http.Request) {
	newArtist := decoderArtist(r.Body)
	artistService.Create(&newArtist)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newArtist)
}

func Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	updatedArtist := artistService.Update(id, decoderArtist(r.Body))

	json.NewEncoder(w).Encode(updatedArtist)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	artistService.Delete(id)

	w.WriteHeader(http.StatusNoContent)
}

func decoderArtist(body io.Reader) models.Artist {
	var newArtist models.Artist
	json.NewDecoder(body).Decode(&newArtist)
	return newArtist
}
