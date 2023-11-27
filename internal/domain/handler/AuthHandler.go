package handler

import (
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

	output := service.Register(user)

	return c.JSON(output)
}

func (r *AuthHandler) Login(c *fiber.Ctx) error {
	LoginRequest := new(request.LoginRequest)
	c.BodyParser(LoginRequest)
	output := service.Login(LoginRequest)

	return c.JSON(output)
}

func (r *AuthHandler) ForgotPassword(c *fiber.Ctx) error {
	req := new(request.ForgotPassword)
	c.BodyParser(req)
	output := service.PostForgotPassword(req)

	return c.JSON(output)
}
