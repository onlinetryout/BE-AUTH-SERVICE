package repository

import (
	"github.com/onlinetryout/BE-AUTH-SERVICE/internal/domain/entities"
	"github.com/onlinetryout/BE-AUTH-SERVICE/internal/domain/request"
)

type AuthInterface interface {
	Register(request *request.RegisterRequest) (entities.User, error)
	GetUserByEmail(user *entities.User, email string) error
}

type AuthRepository struct {
	AuthRepository AuthInterface
}

func NewAuthRepository(authInterface AuthInterface) *AuthRepository {
	return &AuthRepository{
		AuthRepository: authInterface,
	}
}
