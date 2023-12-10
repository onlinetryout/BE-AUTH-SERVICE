package entities

import (
	"time"
)

type ForgotPassword struct {
	ID        uint       `json:"id" gorm:"primaryKey"`
	UserId    uint       `json:"user_id" gorm:"foreignKey:ID"`
	Token     string     `json:"token"`
	SentAt    *time.Time `json:"sent_at"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}
