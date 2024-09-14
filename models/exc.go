package models

import (
	"time"
)

type Exchange struct {
	ID         string    `json:"id" gorm:"primaryKey"`
	User1ID    string    `json:"user1_id"`
	User2ID    string    `json:"user2_id"`
	Comic1ID   string    `json:"comic1_id"`
	Comic2ID   string    `json:"comic2_id"`
	Status     string    `json:"status"` // pending, completed, cancelled
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

