package models

type Review struct {
    ID      string `json:"id" gorm:"primaryKey"`
    ComicID string `json:"comic_id"`
    UserID  string `json:"user_id"`
    Rating  int    `json:"rating"`
    Comment string `json:"comment"`
}

