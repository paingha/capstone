// Copyright 2020 Paingha Joe Alagoa. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package models_test

import(
	"testing"
	"log"
	"bitbucket.com/irb/api/config"
	"bitbucket.com/irb/api/models"
)

func TestCreateUser(t *testing.T){
	if err := config.UseDBTestContext(); err != nil{
		log.Fatal(err)
	}
	var user models.User
	//Add mock user data here
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