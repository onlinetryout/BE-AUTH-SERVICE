package handler

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/onlinetryout/BE-AUTH-SERVICE/internal/domain/request"
	"github.com/onlinetryout/BE-AUTH-SERVICE/internal/domain/response"
	"github.com/onlinetryout/BE-AUTH-SERVICE/internal/domain/service"
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
	output, errValidation := service.Login(LoginRequest)

	if errValidation != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(response.ErrorResponse{
			Success: false,
			Message: "Validation Error",
			Errors:  errValidation,
		})
	}

	return c.JSON(output)
}

func (r *AuthHandler) ForgotPassword(c *fiber.Ctx) error {
	req := new(request.ForgotPassword)
	c.BodyParser(req)
	err := service.PostForgotPassword(req)
	fmt.Println(err)

	return c.JSON(response.SuccessResponse{
		Success: true,
		Message: "Forgot Password Email Sent",
		Data:    nil,
	})
}
