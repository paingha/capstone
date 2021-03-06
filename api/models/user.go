// Copyright 2020 Paingha Joe Alagoa. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package models

import (
	"encoding/base64"
	"fmt"
	"os"
	"time"

	"bitbucket.com/irb/api/config"
	"bitbucket.com/irb/api/security"
	"bitbucket.com/irb/api/stuff"

	jwt "github.com/dgrijalva/jwt-go"
	//Needed for postgres
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

//User - user data struct
type User struct {
	ID                   int        `json:"id,omitempty" sql:"primary_key"`
	IsAdmin              bool       `gorm:"default:false" json:"isAdmin"`
	FirstName            string     `gorm:"not null" json:"firstName"`
	LastName             string     `gorm:"not null" json:"lastName"`
	Email                string     `gorm:"unique;not null" json:"email"`
	PhoneNumber          string     `json:"phoneNumber"`
	Password             string     `gorm:"not null" json:"password"`
	EmailVerified        bool       `gorm:"default:false" json:"emailVerified"`
	VerifyCode           string     `json:"verifyCode"`
	PhoneVerified        bool       `gorm:"default:false" json:"phoneVerified"`
	PhoneVerifyCode      string     `json:"phoneVerifyCode"`
	PhoneVerifySentAt    time.Time  `json:"phone_verify_sent_at"`
	PhoneVerifyExpiresAt time.Time  `json:"phone_verify_expires_at"`
	CreatedAt            time.Time  `json:"created_at"`
	UpdatedAt            time.Time  `json:"updated_at"`
	DeletedAt            *time.Time `json:"deleted_at"`
}

//TableName - table name in database
func (u *User) TableName() string {
	return "users"
}

//GetAllUsers - fetch all users at once
func GetAllUsers(user *[]User, offset int, limit int) (int, error) {
	var count = 0
	if err := config.DB.Model(&User{}).Count(&count).Order("created_at desc").Offset(offset).Limit(limit).Find(user).Error; err != nil {
		return count, err
	}
	return count, nil
}

//CreateUser - create a user
func CreateUser(user *User) (bool, error) {
	var dbUser User
	if err := config.DB.Where("email = ?", user.Email).First(&dbUser).Error; err != nil {
		if err.Error() == "record not found" {
			if errs := config.DB.Create(user).Error; errs != nil {
				return false, errs
			}
			baseURL := os.Getenv("ENV_BASE_URL")
			emailBody := make(map[string]string)
			emailBody["first_name"] = user.FirstName
			emailBody["last_name"] = user.LastName
			emailBody["link"] = fmt.Sprintf("%s/user/0/verify-email?token=%s", baseURL, base64.StdEncoding.EncodeToString([]byte(user.VerifyCode)))
			/*emailInfo := EmailParam{
				To:        user.Email,
				Subject:   "Verify your email",
				BodyParam: emailBody,
				Template:  "TemplateVerifyEmail",
			}
			err := stuff.MailService.Send("email", emailInfo)
			if err != nil {
				fmt.Println(err)
			}*/
			return true, nil
		}
		return false, err
	}
	return false, nil
}

//LoginUser - fetch one user
func LoginUser(user *User) (User, string, error) {
	var dbUser User
	jwtSecretByte := []byte(os.Getenv("JWT_SECRET"))
	expiresAt := time.Now().Add(1200 * time.Minute)
	if err := config.DB.Model(&user).Where(&User{Email: user.Email}).First(&dbUser).Error; err != nil {
		return User{}, "", err
	}
	//compare db password hash and password provided
	resp := security.VerifyHash([]byte(dbUser.Password), []byte(user.Password))
	if !resp {
		return User{}, "", nil
	}
	fmt.Println("--------------------------------------")
	fmt.Println(dbUser.IsAdmin)
	fmt.Println("--------------------------------------")
	claims := &security.Claims{
		UserID:  dbUser.ID,
		IsAdmin: dbUser.IsAdmin,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expiresAt.Unix(),
		},
	}
	// Declare the token with the algorithm used for signing, and the claims
	tokens := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	tokenString, errs := tokens.SignedString(jwtSecretByte)
	if errs != nil {
		return User{}, "", errs
	}
	return dbUser, tokenString, nil

}

//GetUser - fetch one user
func GetUser(user *User, id int) error {
	if err := config.DB.Where("id = ?", id).First(user).Error; err != nil {
		return err
	}
	return nil
}

//VerifyEmailUser - verify user's email
func VerifyEmailUser(user *User, token string) error {
	if err := config.DB.Model(&user).Where(&User{VerifyCode: token}).Updates(map[string]interface{}{"email_verified": true, "verify_code": ""}).Error; err != nil {
		return err
	}
	return nil
}

//SendVerifyPhoneUser - send verification code to user's phone number
func SendVerifyPhoneUser(user *User, id int, code string, medium string) error {
	current := time.Now()
	future := current.Add(time.Minute * 30) //expires after 30 minutes of being sent
	if err := config.DB.Model(&user).Where(&User{ID: id}).Updates(map[string]interface{}{"phone_number": user.PhoneNumber, "phone_verify_sent_at": current, "phone_verify_expires_at": future, "phone_verify_code": code}).Error; err != nil {
		return err
	}
	message := Message{
		Content: "NSU IRB Verification Code: " + code,
		To:      user.PhoneNumber,
		Medium:  medium,
	}
	errs := stuff.SmsService.Send("sms", message) //Sends to sms message queue
	if errs != nil {
		return errs
	}
	return nil
}

//VerifyPhoneUser - verifies the verify code and expiry time and then sets phone_verified to true
func VerifyPhoneUser(user *User, id int, token string) (bool, error) {
	var dbUser User
	current := time.Now()
	if err := config.DB.Where("id = ?", id).First(&dbUser).Error; err != nil {
		return false, err
	}
	if current.Before(dbUser.PhoneVerifyExpiresAt) && token == dbUser.PhoneVerifyCode {
		if errs := config.DB.Model(&user).Where(&User{PhoneVerifyCode: token}).Updates(map[string]interface{}{"phone_verified": true, "phone_verify_code": ""}).Error; errs != nil {
			return false, errs
		}
		return true, nil
	}
	return false, nil
}

//UpdateUser - update a user
func UpdateUser(user *User, id int) error {
	if err := config.DB.Model(&user).Omit("is_admin", "email_verified", "password", "verify_code", "phone_verified", "phone_verify_code", "created_at", "updated_at", "deleted_at", "phone_verify_sent_at", "phone_verify_expires_at").Updates(user).Error; err != nil {
		return err
	}
	return nil
}

//DeleteUser - delete a user
func DeleteUser(id int) error {
	if err := config.DB.Where("id = ?", id).Unscoped().Delete(User{}).Error; err != nil {
		return err
	}
	return nil
}

//ForgotUser - sends a forgot password email to a user
func ForgotUser(user *User) (bool, error) {
	var dbUser User
	if err := config.DB.Where("email = ?", user.Email).First(&dbUser).Error; err != nil {
		return false, err
	}
	//Send Verification email to Rabbitmq
	jwtSecretByte := []byte(os.Getenv("JWT_SECRET"))
	expiresAt := time.Now().Add(30 * time.Minute)
	emailBody := make(map[string]string)
	emailBody["first_name"] = dbUser.FirstName
	emailBody["last_name"] = dbUser.LastName
	//Create a jwt token here
	claims := &security.Claims{
		UserID: dbUser.ID,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expiresAt.Unix(),
		},
	}
	// Declare the token with the algorithm used for signing, and the claims
	tokens := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, errs := tokens.SignedString(jwtSecretByte)
	if errs != nil {
		return false, errs
	}
	baseURL := os.Getenv("ENV_BASE_URL")
	emailBody["link"] = fmt.Sprintf("%s/user/0/forgot-password?token=%s", baseURL, base64.StdEncoding.EncodeToString([]byte(tokenString)))
	emailInfo := EmailParam{
		To:        user.Email,
		Subject:   "Password Reset",
		BodyParam: emailBody,
		Template:  "TemplateResetEmail",
	}
	errss := stuff.MailService.Send("email", emailInfo)
	if errss != nil {
		fmt.Println(errss)
		return false, errss
	}
	return true, nil
}
