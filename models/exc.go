package models

type ExchangeStatus string

const (
	Pending  ExchangeStatus = "pending"
	Accepted ExchangeStatus = "accepted"
	Declined ExchangeStatus = "declined"
)

type Exchange struct {
	ID          string         `json:"id" gorm:"primaryKey"`
	ComicID     string         `json:"comic_id" gorm:"not null"`
	RequesterID string         `json:"requester_id" gorm:"not null"`
	OwnerID     string         `json:"owner_id" gorm:"not null"`
	Status      ExchangeStatus `json:"status" gorm:"default:'pending'"`
}
