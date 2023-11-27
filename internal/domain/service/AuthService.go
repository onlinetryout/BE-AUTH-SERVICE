package service

import (
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/onlinetryout/BE-AUTH-SERVICE/internal/config"
	"github.com/onlinetryout/BE-AUTH-SERVICE/internal/domain/entities"
	"github.com/onlinetryout/BE-AUTH-SERVICE/internal/domain/repository"
	"github.com/onlinetryout/BE-AUTH-SERVICE/internal/domain/request"
	"github.com/onlinetryout/BE-AUTH-SERVICE/internal/domain/response"
	"github.com/onlinetryout/BE-AUTH-SERVICE/internal/infra/database"
	"github.com/onlinetryout/BE-AUTH-SERVICE/pkg/utils"
	"log"
)

func Register(req *request.RegisterRequest) (entities.User, []request.ErrorResponse) {
	validate := validator.New()
	errs := validate.Struct(req)
	validationErrors := []request.ErrorResponse{}
	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			var elem request.ErrorResponse
			elem.FailedField = err.Field() // Export struct field name
			elem.Tag = err.Tag()           // Export struct tag
			elem.Value = err.Value()       // Export field value
			validationErrors = append(validationErrors, elem)
		}
		return entities.User{}, validationErrors
	}

	//CHeck email unique
	var userCount int64
	database.DB.Model(&entities.User{}).Where("email", req.Email).Count(&userCount)
	if userCount > 0 {
		var elem request.ErrorResponse
		elem.FailedField = "Email"
		elem.Tag = "Email already used"
		elem.Value = ""
		validationErrors = append(validationErrors, elem)
		return entities.User{}, validationErrors
	}

	repo := repository.NewAuthRepository(&repository.AuthMysql{})

	NewUser, err := repo.AuthRepository.Register(req)
	if err != nil {
		return entities.User{}, nil
	}
	return NewUser, nil
}

func Login(req *request.LoginRequest) (interface{}, []request.ErrorResponse) {
	repo := repository.NewAuthRepository(&repository.AuthMysql{})

	validate := validator.New()
	errs := validate.Struct(req)
	validationErrors := []request.ErrorResponse{}
	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			var elem request.ErrorResponse
			elem.FailedField = err.Field() // Export struct field name
			elem.Tag = err.Tag()           // Export struct tag
			elem.Value = err.Value()       // Export field value
			validationErrors = append(validationErrors, elem)
		}
		return "", validationErrors
	}

	//Validation email and password
	var user entities.User
	//email checking
	if err := repo.AuthRepository.GetUserByEmail(&user, req); err != nil {
		var elem request.ErrorResponse
		elem.FailedField = "Email"
		elem.Tag = "Email not found"
		elem.Value = ""
		validationErrors = append(validationErrors, elem)
		return "", validationErrors
	}
	//Check Password
	if isValid := utils.CheckPasswordHash(req.Password, user.Password); !isValid {
		var elem request.ErrorResponse
		elem.FailedField = "Password"
		elem.Tag = "Wrong Password"
		elem.Value = ""
		validationErrors = append(validationErrors, elem)
		return "", validationErrors
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
	return output, nil
}

func PostForgotPassword(req *request.ForgotPassword) error {
	return nil
}
