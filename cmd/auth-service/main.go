package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/onlinetryout/BE-AUTH-SERVICE/internal/config"
	"github.com/onlinetryout/BE-AUTH-SERVICE/internal/infra/database"
	"github.com/onlinetryout/BE-AUTH-SERVICE/internal/infra/migration"
	"github.com/onlinetryout/BE-AUTH-SERVICE/route"
)

func main() {
	//Init Config
	config.ConfigInit()
	//Init Database
	database.DatabaseInit()
	migration.RunMigration()

	// Fiber instance
	app := fiber.New()

	//Init Route
	route.RouteInit(app)

	// start server
	err := app.Listen(":" + config.ConfigApp.Port)
	if err != nil {
		fmt.Println(err)
	}

}
