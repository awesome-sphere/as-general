package models

type Movie struct {
	Id          int    `json:"id" gorm:"autoincrement;primaryKey"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Duration    int    `json:"duration"`
	Poster      string `json:"poster"`
	Trailer     string `json:"trailer"`
}
