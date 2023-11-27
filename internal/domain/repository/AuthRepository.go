package repository

import (
	"github.com/google/uuid"
	"github.com/onlinetryout/BE-AUTH-SERVICE/internal/domain/entities"
	"github.com/onlinetryout/BE-AUTH-SERVICE/internal/domain/request"
	"github.com/onlinetryout/BE-AUTH-SERVICE/internal/infra/database"
	"github.com/onlinetryout/BE-AUTH-SERVICE/pkg/utils"
	"log"
)

type AuthMysql struct {
}

func (a AuthMysql) Register(request *request.RegisterRequest) (entities.User, error) {

	hashedPassword, err := utils.HashingPassword(request.Password)

	if err != nil {
		return entities.User{}, err
	}

	newUser := entities.User{
		Uuid:     uuid.New(),
		Name:     request.Name,
		Password: hashedPassword,
		Email:    request.Email,
		Address:  request.Address,
		Phone:    request.Phone,
	}

	err = database.DB.Create(&newUser).Error

	if err != nil {
		log.Println("error auth repository: ", err)
		return entities.User{}, err
	}

	return newUser, nil
}

func (a AuthMysql) GetUserByEmail(user *entities.User, req *request.LoginRequest) error {
	err := database.DB.Where("email", req.Email).First(&user).Error
	if err != nil {
		return err
	}
	return nil
}
