package models

type Song struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Duration  int    `json:"duration"`
	Album     string `json:"album"`
	Artist_id int    `json:"artistId"`
}
