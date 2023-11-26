package service

import (
	"github.com/go-playground/validator/v10"
	"github.com/onlinetryout/BE-AUTH-SERVICE/internal/domain/entities"
	"github.com/onlinetryout/BE-AUTH-SERVICE/internal/domain/repository"
	"github.com/onlinetryout/BE-AUTH-SERVICE/internal/domain/request"
	"github.com/onlinetryout/BE-AUTH-SERVICE/internal/infra/database"
)

func Register(req *request.RegisterRequest) (bool, []request.ErrorResponse) {
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
		return false, validationErrors
	}

	//CHeck email unique
	var userCount int64
	query := database.DB.Model(&entities.User{}).Where("email", req.Email)
	query.Count(&userCount)
	if userCount > 0 {
		var elem request.ErrorResponse
		elem.FailedField = "Email"
		elem.Tag = "Email already used"
		elem.Value = ""
		validationErrors = append(validationErrors, elem)
		return false, validationErrors
	}

	repo := repository.NewAuthRepository(&repository.AuthMysql{})

	if _, err := repo.AuthRepository.Register(req); err != nil {
		return false, nil
	}
	return true, validationErrors
}
