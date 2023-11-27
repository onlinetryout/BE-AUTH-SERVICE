package repository

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/onlinetryout/BE-AUTH-SERVICE/internal/config"
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

func (a AuthMysql) Login(user entities.User) (string, error) {
	claims := jwt.MapClaims{}
	claims["id"] = user.ID
	claims["name"] = user.Name
	claims["email"] = user.Email
	claims["exp"] = config.ConfigJwt.Expired.Unix()
	token, err := utils.GenerateToken(&claims)
	if err != nil {
		log.Println("error on generate token", err)
	}

	return token, nil
}
