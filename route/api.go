package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/onlinetryout/BE-AUTH-SERVICE/internal/domain/handler"
)

func RouteInit(route *fiber.App) {
	api := route.Group("/api")
	v1 := api.Group("/v1")

	//Route List
	AuthHandler := handler.NewAuthHandler()
	v1.Post("/register", AuthHandler.Register)
	v1.Post("/login", AuthHandler.Login)
	v1.Post("/forgot-password", AuthHandler.ForgotPassword)
}
