package migration

import (
	"fmt"
	"github.com/onlinetryout/BE-AUTH-SERVICE/internal/domain/entities"
	"github.com/onlinetryout/BE-AUTH-SERVICE/internal/infra/database"
	"log"
)

func RunMigration() {
	err := database.DB.AutoMigrate(&entities.User{}, &entities.ForgotPassword{})
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Database Migrated Successfully")
}
