// Copyright 2020 Paingha Joe Alagoa. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package models_test

import(
	"testing"
	"log"
	"encoding/base64"
	"bitbucket.com/irb/api/config"
	"bitbucket.com/irb/api/utils"
	"bitbucket.com/irb/api/models"
	"github.com/brianvoe/gofakeit/v5"
)

func TestCreateUser(t *testing.T){
	if err := config.UseDBTestContext(); err != nil{
		log.Fatal(err)
	}
	gofakeit.Seed(0)
	user := models.User{
		FirstName: gofakeit.FirstName(),
		LastName: gofakeit.LastName(),
		Email: gofakeit.Email(),
		PhoneNumber: gofakeit.Phone(),
		Password: gofakeit.Password(true, true, true, true, true, 6),
	}
	status, err := models.CreateUser(&user)
	if err != nil {
		log.Print(err)
	} else {
		if status != true {
			log.Print("Account already exists")
		} else {
			log.Print("Account created successfully")
		}
	}
}

func TestLoginUser(t *testing.T){
	if err := config.UseDBTestContext(); err != nil{
		log.Fatal(err)
	}
	var user models.User
	//Add mock user data here
	user, token, err := models.LoginUser(&user)
	if err != nil {
		log.Print(err)
	} else {
		log.Println(user)
		log.Println(token)
	}
}

func TestGetUser(t *testing.T){
	if err := config.UseDBTestContext(); err != nil{
		log.Fatal(err)
	}
	var user models.User
	var userID int
	err := models.GetUser(&user, userID)
	if err != nil {
		log.Print(err)
	} else {
		log.Println(user)
	}
}

func TestGetUsers(t *testing.T){
	if err := config.UseDBTestContext(); err != nil{
		log.Fatal(err)
	}
	var users []models.User
	var count, offset, limit int = 0, 0, 6
	count, err := models.GetAllUsers(&users, offset, limit)
	if err != nil {
		log.Print(err)
	} else {
		log.Println(users)
		log.Println(count)
	}
}

func TestSendVerifyPhoneUser(t *testing.T){
	if err := config.UseDBTestContext(); err != nil{
		log.Fatal(err)
	}
	code := utils.GenerateRandomInt(5)
	medium := "whatsapp"
	var user models.User
	var id int;
	err := models.SendVerifyPhoneUser(&user, id, code, medium)
	if err != nil {
		log.Print(err)
	} else {
		log.Println("Code sent successfully")
	}
}

func TestForgotUser(t *testing.T){
	var user models.User
	if err := config.UseDBTestContext(); err != nil{
		log.Fatal(err)
	}
	resp, err := models.ForgotUser(&user)
	if err != nil{
		log.Print(err)
	}else{
		if !resp && err == nil {
			log.Println("Could not send the reset email")
		} else {
			log.Println("Password Reset Email sent successfully")
		}
	}
}

func TestDeleteUser(t *testing.T){
	var id int = 0;
	if err := config.UseDBTestContext(); err != nil{
		log.Fatal(err)
	}
	err := models.DeleteUser(id)
	if err != nil{
		log.Print(err)
	}else{
		log.Println("User Deleted successfully")
	}
}

func TestUpdateUser(t *testing.T){
	var user models.User
	var id int = 0;
	if err := config.UseDBTestContext(); err != nil{
		log.Fatal(err)
	}
	err := models.UpdateUser(&user, id)
	if err != nil{
		log.Print(err)
	}else{
		log.Println("Password Reset Email sent successfully")
	}
}

func TestVerifyEmailUser(t *testing.T){
	var token string = ""
	var user models.User
	if err := config.UseDBTestContext(); err != nil{
		log.Fatal(err)
	}
	if tokenDecoded, err := base64.StdEncoding.DecodeString(token); err != nil {
		log.Println("Email verify token decode error")
	} else {
		log.Println(string(tokenDecoded))
		token = string(tokenDecoded)
	}
	err := models.VerifyEmailUser(&user, token)
	if err != nil{
		log.Print(err)
	}else{
		log.Println("Password Reset Email sent successfully")
	}
}

func TestVerifyPhoneCodeUser(t *testing.T){
	var user models.User
	var token string = ""
	var id int = 0;
	if err := config.UseDBTestContext(); err != nil{
		log.Fatal(err)
	}
	state, errs := models.VerifyPhoneUser(&user, id, token)
	if errs != nil {
		log.Println("An error occured while verifying user's phone number")
		log.Println(errs)
	} else {
		if state == true && errs == nil {
			log.Println("Phone verifed successfully")
		} else {
			log.Println("Token expired or invalid")
		}
	}
}