package repository

import (
	"errors"
	"fmt"
	"github.com/go-mail/mail"
	"github.com/google/uuid"
	"github.com/onlinetryout/BE-AUTH-SERVICE/internal/config"
	"github.com/onlinetryout/BE-AUTH-SERVICE/internal/domain/entities"
	"github.com/onlinetryout/BE-AUTH-SERVICE/internal/domain/request"
	"github.com/onlinetryout/BE-AUTH-SERVICE/internal/infra/database"
	"github.com/onlinetryout/BE-AUTH-SERVICE/pkg/utils"
	"log"
	"strconv"
	"time"
)

type AuthMysql struct {
}

func (a AuthMysql) Register(request *request.RegisterRequest) (entities.User, error) {

	hashedPassword, err := utils.HashingPassword(request.Password)

	if err != nil {
		return entities.User{}, err
	}

	newUser := entities.User{
		Uuid:     uuid.New(),
		Name:     request.Name,
		Password: hashedPassword,
		Email:    request.Email,
		Address:  request.Address,
		Phone:    request.Phone,
	}

	err = database.DB.Create(&newUser).Error

	if err != nil {
		log.Println("error auth repository: ", err)
		return entities.User{}, err
	}

	return newUser, nil
}

func (a AuthMysql) GetUserByEmail(user *entities.User, email string) error {
	err := database.DB.Where("email", email).First(&user).Error
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (a AuthMysql) SendEmailForgotPassword(user entities.User) error {
	var userForgot entities.ForgotPassword
	database.DB.Where("user_id", user.ID).First(&userForgot)
	m := mail.NewMessage()
	m.SetHeader("From", config.ConfigEmail.From)

	m.SetHeader("To", user.Email)

	m.SetHeader("Subject", "Forgot Password Email!")
	urlForgot := config.ConfigForgotPassword.FrontUrl + "/?token=" + userForgot.Token
	fullURL := "https://" + urlForgot // Include the full URL with the scheme

	emailBody := "Klik <a href=\"" + fullURL + "\">disini</a> untuk mengganti password!"

	m.SetBody("text/html", emailBody)

	m.SetBody("text/html", emailBody)

	//m.Attach("lolcat.jpg")

	d := mail.NewDialer(config.ConfigEmail.Host, config.ConfigEmail.Port, config.ConfigEmail.Username, config.ConfigEmail.Password)

	// Send the email to Kate, Noah and Oliver.

	if err := d.DialAndSend(m); err != nil {
		log.Println(err)
	}

	//Change Sent At
	currentTime := time.Now()
	userForgot.SentAt = &currentTime
	database.DB.Updates(&userForgot)

	return nil
}

func (a AuthMysql) GenerateTokenForgotPassword(user entities.User) error {
	var err error
	err = nil
	var userToken entities.ForgotPassword
	database.DB.Where("user_id", user.ID).First(&userToken)
	if userToken.ID != 0 {
		//check last sent email
		expiredMinute := config.ConfigForgotPassword.ExpiredTime
		expirationDuration := time.Duration(expiredMinute) * time.Minute
		// Assuming userToken.SentAt is a time.Time field
		if userToken.SentAt != nil {
			if !userToken.SentAt.IsZero() {
				currentTime := time.Now()
				expirationTime := userToken.SentAt.Add(expirationDuration)
				if currentTime.After(expirationTime) {
					// Can email again
					//Delete previous data
					database.DB.Delete(&userToken)
					userToken.UserId = user.ID
					userToken.Token = utils.GenerateRandomString(255)
					userToken.SentAt = nil
					err = database.DB.Create(&userToken).Error
				} else {
					//Cannot send email
					expiredTimeStr := strconv.Itoa(int(config.ConfigForgotPassword.ExpiredTime))

					err = errors.New("Tidak dapat mengirim email kembali, harap tunggu dalam " + expiredTimeStr + " menit")
				}
			} else {
				//Can Send email
				database.DB.Delete(&userToken)
				userToken.UserId = user.ID
				userToken.Token = utils.GenerateRandomString(255)
				userToken.SentAt = nil
				err = database.DB.Create(&userToken).Error
			}
		}
	} else {
		userToken.UserId = user.ID
		userToken.Token = utils.GenerateRandomString(255)
		userToken.SentAt = nil
		err = database.DB.Create(&userToken).Error
	}
	return err
}

func (a AuthMysql) UserFromToken(token string, user *entities.User) error {
	var forgotPassword entities.ForgotPassword
	err := database.DB.Where("token = ?", token).First(&forgotPassword).Error
	if err != nil {
		fmt.Println("Error querying database:", err)
		var errs error
		errs = errors.New("Token Not Found")
		return errs
	}
	//Check Token still usable
	// Check if the token is still usable (e.g., within 30 minutes)
	tokenExpiration := forgotPassword.SentAt.Add(30 * time.Minute)
	currentTime := time.Now()

	if currentTime.After(tokenExpiration) {
		var errs error
		errs = errors.New("Token Expired")
		return errs
	}
	//end check

	err = database.DB.Where("id = ?", forgotPassword.UserId).First(&user).Error
	if err != nil {
		fmt.Println(forgotPassword.UserId)
		fmt.Println(err)
		var errs error
		errs = err
		return errs
	}

	return nil
}

func (a AuthMysql) ChangePassword(user entities.User, newPassword string) error {
	err := database.DB.Model(&user).Update("Password", newPassword).Error
	if err != nil {
		// Handle the error, e.g., log it or return it
		fmt.Println("Error updating password in the database:", err)
		return errors.New("failed to update password")
	}
	return nil
}
