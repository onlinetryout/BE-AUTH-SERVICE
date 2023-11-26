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
	"log"
)

func (r *AuthHandler) Register(c *fiber.Ctx) error {

	//Mapping Register Request
	user := new(request.RegisterRequest)
	if err := c.BodyParser(user); err != nil {
		// Handle parsing error
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Error parsing request body",
		})
	}

	success, validationErrors := service.Register(user)
	if !success {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Validation Error",
			"errors":  validationErrors,
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "Register successfully",
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

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Validation error",
			"error":   validationErrors,
		})
	}
	var user entities.User
	//Check Email
	if err := database.DB.Where("email", LoginRequest.Email).First(&user).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "failed",
			"message": "email not found",
		})
	}
	//Check Password
	isValid := utils.CheckPasswordHash(LoginRequest.Password, user.Password)

	if !isValid {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "failed",
			"message": "wrong password",
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
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "failed",
			"message": "Internal server error",
			"error":   err,
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
		User:      UserResponse,
	}
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Login success",
		"data":    output,
	})
}
