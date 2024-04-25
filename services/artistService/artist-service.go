package artistService

import (
	"github.com/gabrielsilper/golang-api-rest/database"
	"github.com/gabrielsilper/golang-api-rest/models"
)

func FindAll() []models.Artist {
	var artists []models.Artist
	database.DB.Find(&artists)
	return artists
}

func FindById(id string) models.Artist {
	var artist models.Artist
	database.DB.First(&artist, id)
	return artist
}
