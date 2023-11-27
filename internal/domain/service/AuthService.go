package service

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/onlinetryout/BE-AUTH-SERVICE/internal/config"
	"github.com/onlinetryout/BE-AUTH-SERVICE/internal/domain/entities"
	"github.com/onlinetryout/BE-AUTH-SERVICE/internal/domain/repository"
	"github.com/onlinetryout/BE-AUTH-SERVICE/internal/domain/request"
	"github.com/onlinetryout/BE-AUTH-SERVICE/internal/domain/response"
	"github.com/onlinetryout/BE-AUTH-SERVICE/pkg/utils"
	"log"
)

func Register(req *request.RegisterRequest) interface{} {
	repo := repository.NewAuthRepository(&repository.AuthMysql{})
	pass, validationMessage := utils.Validate(req)
	if !pass {
		res := response.ErrorResponse{
			Success: false,
			Message: "Validation Error",
			Errors:  validationMessage,
		}
		return res
	}

	//CHeck email unique
	var validationErrors []request.ErrorResponse
	var user entities.User
	err := repo.AuthRepository.GetUserByEmail(&user, req.Email)
	if err != nil {
		var elem request.ErrorResponse
		elem.FailedField = "Email"
		elem.Tag = "Email already used"
		elem.Value = ""
		validationErrors = append(validationErrors, elem)
		res := response.ErrorResponse{
			Success: false,
			Message: "Validation Error",
			Errors:  validationMessage,
		}
		return res
	}

	NewUser, err := repo.AuthRepository.Register(req)
	if err != nil {
		return nil
	}
	return response.SuccessResponse{
		Success: true,
		Message: "Register success",
		Data:    NewUser,
	}
}

func Login(req *request.LoginRequest) interface{} {
	repo := repository.NewAuthRepository(&repository.AuthMysql{})

	isPass, validationMessage := utils.Validate(req)
	if !isPass {
		res := response.ErrorResponse{
			Success: false,
			Message: "Validation Error",
			Errors:  validationMessage,
		}
		return res
	}

	//Validation email and password
	var user entities.User
	var validationErrors []request.ErrorResponse
	//email checking
	if err := repo.AuthRepository.GetUserByEmail(&user, req.Email); err != nil {
		var elem request.ErrorResponse
		elem.FailedField = "Email"
		elem.Tag = "Email not found"
		elem.Value = ""
		validationErrors = append(validationErrors, elem)
		res := response.ErrorResponse{
			Success: false,
			Message: "Validation Error",
			Errors:  validationErrors,
		}
		return res
	}
	//Check Password
	if isValid := utils.CheckPasswordHash(req.Password, user.Password); !isValid {
		var elem request.ErrorResponse
		elem.FailedField = "Password"
		elem.Tag = "Wrong Password"
		elem.Value = ""
		validationErrors = append(validationErrors, elem)
		res := response.ErrorResponse{
			Success: false,
			Message: "Validation Error",
			Errors:  validationErrors,
		}
		return res
	}

	claims := jwt.MapClaims{}
	claims["id"] = user.ID
	claims["name"] = user.Name
	claims["email"] = user.Email
	claims["exp"] = config.ConfigJwt.Expired.Unix()
	token, err := utils.GenerateToken(&claims)

	if err != nil {
		log.Println(err)
	}
	userResponse := response.UserResponse{
		Uuid:      user.Uuid,
		Name:      user.Name,
		Email:     user.Email,
		Address:   user.Address,
		Phone:     user.Phone,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.CreatedAt,
	}
	output := response.LoginResponse{
		TokenType: "Bearer",
		Token:     token,
		Data:      userResponse,
	}
	return output
}

func PostForgotPassword(req *request.ForgotPassword) interface{} {
	pass, validationMessage := utils.Validate(req)
	if !pass {
		res := response.ErrorResponse{
			Success: false,
			Message: "Validation Error",
			Errors:  validationMessage,
		}
		return res
	}
	return nil
}
