package models

import "time"

type Review struct {
	ID        string    `json:"id" gorm:"primaryKey"`
	UserID    string    `json:"user_id"`
	ComicID   string    `json:"comic_id"`
	Rating    int       `json:"rating"` // rating de 1 a 5
	Comment   string    `json:"comment"`
	CreatedAt time.Time `json:"created_at"`
}

