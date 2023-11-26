package handler

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/onlinetryout/BE-AUTH-SERVICE/internal/config"
	"github.com/onlinetryout/BE-AUTH-SERVICE/internal/domain/entities"
	"github.com/onlinetryout/BE-AUTH-SERVICE/internal/domain/request"
	"github.com/onlinetryout/BE-AUTH-SERVICE/internal/domain/response"
	"github.com/onlinetryout/BE-AUTH-SERVICE/internal/domain/service"
	"github.com/onlinetryout/BE-AUTH-SERVICE/internal/infra/database"
	"github.com/onlinetryout/BE-AUTH-SERVICE/pkg/utils"
)

func (r *AuthHandler) Register(c *fiber.Ctx) error {

	//Mapping Register Request
	user := new(request.RegisterRequest)
	if err := c.BodyParser(user); err != nil {
		// Handle parsing error
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Success: false,
			Message: "Error Parsing Request Data",
		})
	}

	newUser, validationErrors := service.Register(user)

	if validationErrors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Success: false,
			Message: "Validation Error",
			Errors:  validationErrors,
		})
	}

	return c.JSON(response.SuccessResponse{
		Success: true,
		Message: "Register user succesfully",
		Data:    newUser,
	})
}

func (r *AuthHandler) Login(c *fiber.Ctx) error {
	LoginRequest := new(request.LoginRequest)
	c.BodyParser(LoginRequest)
	validate := validator.New()

	errs := validate.Struct(LoginRequest)
	if errs != nil {
		validationErrors := []request.ErrorResponse{}
		for _, err := range errs.(validator.ValidationErrors) {
			var elem request.ErrorResponse
			elem.FailedField = err.Field() // Export struct field name
			elem.Tag = err.Tag()           // Export struct tag
			elem.Value = err.Value()       // Export field value
			validationErrors = append(validationErrors, elem)
		}

		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Success: false,
			Message: "Validation Error",
			Errors:  validationErrors,
		})
	}
	var user entities.User
	//Check Email
	if err := database.DB.Where("email", LoginRequest.Email).First(&user).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(response.ErrorResponse{
			Success: false,
			Message: "Email Not Found",
		})
	}
	//Check Password
	isValid := utils.CheckPasswordHash(LoginRequest.Password, user.Password)

	if !isValid {
		return c.Status(fiber.StatusNotFound).JSON(response.ErrorResponse{
			Success: false,
			Message: "Wrong Password",
		})
	}

	//Generate JWT
	claims := jwt.MapClaims{}
	claims["id"] = user.ID
	claims["name"] = user.Name
	claims["email"] = user.Email
	claims["exp"] = config.ConfigJwt.Expired.Unix()

	token, err := utils.GenerateToken(&claims)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse{
			Success: false,
			Message: "Internal server error",
			Errors:  err,
		})
	}

	UserResponse := response.UserResponse{
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
		Data:      UserResponse,
	}
	return c.JSON(response.SuccessResponse{
		Success: true,
		Message: "Login Success",
		Data:    output,
	})
}
