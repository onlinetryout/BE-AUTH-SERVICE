package response

import (
	"github.com/google/uuid"
	"time"
)

type LoginResponse struct {
	TokenType string
	Token     string
	User      UserResponse
}

type UserResponse struct {
	Uuid      uuid.UUID `json:"uuid"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Address   string    `json:"address"`
	Phone     string    `json:"phone"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
