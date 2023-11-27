package response

import (
	"github.com/google/uuid"
	"time"
)

type LoginResponse struct {
	TokenType string       `json:"token_type"`
	Token     string       `json:"token"`
	Data      UserResponse `json:"data"`
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
