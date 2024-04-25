package artistController

import (
	"encoding/json"
	"net/http"

	"github.com/gabrielsilper/golang-api-rest/models"
)

func GetAll(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.ArtistSlice)
}
