package models

type Comic struct {
	ID        string `json:"id" gorm:"primaryKey"`
	Title     string `json:"title"`
	Author    string `json:"author"`
	Publisher string `json:"publisher"`
	CoverURL  string `json:"cover_url"`
	Year      int    `json:"year"`
	Owner     User   `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	UserID    string `json:"user_id"`
}

type Comics []Comic
