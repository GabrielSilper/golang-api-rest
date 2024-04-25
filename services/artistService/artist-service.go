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

func Create(newArtist *models.Artist) {
	database.DB.Create(newArtist)
}

func Delete(id string) {
	database.DB.Delete(&models.Artist{}, id)
}

func Update(id string, updatedArtist models.Artist) models.Artist {
	artist := FindById(id)
	artist.Name = updatedArtist.Name
	artist.Nationality = updatedArtist.Nationality
	database.DB.Save(&artist)
	return artist
}
