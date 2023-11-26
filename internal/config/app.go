package config

import "time"

type MySql struct {
	Host     string
	Port     string
	Database string
	Username string
	Password string
}

type App struct {
	Port    string
	Appname string
}

type Jwt struct {
	Key     string
	Expired time.Time
}

var ConfigJwt *Jwt
var ConfigMysql *MySql
var ConfigApp *App

func ConfigInit() {
	ConfigMysql = &MySql{
		Host:     "127.0.0.1",
		Port:     "3306",
		Database: "onlinetryout",
		Username: "root",
		Password: "mypassword",
	}

	ConfigApp = &App{
		Port:    "8080",
		Appname: "BE-SERVICE",
	}

	ConfigJwt = &Jwt{
		Key:     "jdsklafjie9wpasf019fa83032lfmp0283jasnf083743",
		Expired: time.Now().Add(time.Minute * 360), //minutes
	}
}
