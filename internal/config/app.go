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

type Email struct {
	From     string
	Host     string
	Port     int
	Username string
	Password string
}

type ForgotPassword struct {
	FrontUrl    string
	ExpiredTime int
}

var ConfigJwt *Jwt
var ConfigMysql *MySql
var ConfigApp *App
var ConfigEmail *Email
var ConfigForgotPassword *ForgotPassword

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

	ConfigEmail = &Email{
		From:     "forgotpassword@warungacehbangari.com",
		Host:     "smtp.hostinger.com",
		Port:     587, //465
		Username: "forgotpassword@warungacehbangari.com",
		Password: "qweasdzxc123!Q",
	}

	ConfigForgotPassword = &ForgotPassword{
		FrontUrl:    "127.0.0.1:8000/url",
		ExpiredTime: 10, //in minutes
	}

}
