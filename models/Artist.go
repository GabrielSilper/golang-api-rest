package models

type Artist struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Nationality string `json:"nationality"`
}

var ArtistSlice = []Artist{
	{1, "Jimi Hendrix", "American"},
	{2, "Adele", "British"},
	{3, "Tupac Shakur", "American"},
	{4, "Wolfgang Amadeus Mozart", "Austrian"},
}
