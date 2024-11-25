package models

type Comic struct {
	ID          string `json:"id" gorm:"primaryKey"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	UserID      string `json:"user_id"`
	Publisher   string `json:"publisher"`
	Description string `json:"description"`
	Genre       string `json:"genre"`
	CoverURL    string `json:"cover_url"`
	Pages       int    `json:"pages"`
	Edition     string `json:"edition"`
	Year        int    `json:"year"`
}

type Comics []Comic
