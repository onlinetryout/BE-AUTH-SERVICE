package utils

import (
	"github.com/go-playground/validator/v10"
	"github.com/onlinetryout/BE-AUTH-SERVICE/internal/domain/request"
)

func Validate(req interface{}) (pass bool, validationMessage []request.ErrorResponse) {
	pass = true

	validate := validator.New()
	errs := validate.Struct(req)
	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			var elem request.ErrorResponse
			elem.FailedField = err.Field()
			elem.Tag = err.Tag()
			elem.Value = err.Value()
			validationMessage = append(validationMessage, elem)
		}
		pass = false
	}

	return pass, validationMessage
}