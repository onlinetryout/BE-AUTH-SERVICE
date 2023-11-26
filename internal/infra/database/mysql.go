package database

import (
	"fmt"
	"github.com/onlinetryout/BE-AUTH-SERVICE/internal/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DatabaseInit() {
	var err error
	MYSQL := config.ConfigMysql.Username + ":" + config.ConfigMysql.Password + "@tcp(" + config.ConfigMysql.Host + ":" + config.ConfigMysql.Port + ")/" + config.ConfigMysql.Database + "?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := MYSQL

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Cannot connect to database")
	}

	fmt.Println("Connected to database")
}
