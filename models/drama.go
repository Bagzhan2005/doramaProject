package models

import "time"

type Drama struct {
	ID            uint      `json:"id" gorm:"primaryKey"`
	Title         string    `json:"title"`
	OriginalTitle string    `json:"original_title"`
	Country       string    `json:"country"`
	ReleaseYear   int       `json:"release_year"`
	Episodes      int       `json:"episodes"`
	Status        string    `json:"status"`
	Description   string    `json:"description"`
	Rating        float64   `json:"rating"`
	CreatedAt     time.Time `json:"created_at"`
	Genres        []Genre   `json:"genres" gorm:"many2many:drama_genres;"`
}

type Genre struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}
