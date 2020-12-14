// Copyright 2020 Paingha Joe Alagoa. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"testing"

	"bitbucket.com/irb/api/config"
	"bitbucket.com/irb/api/models"
)

func TestSendMail(t *testing.T) {
	err := config.InitConfig(&cfg)
	if err != nil {
		t.Error(err)
	}

	emailBody := make(map[string]string)
	emailBody["first_name"] = "user.FirstName"
	emailBody["last_name"] = "user.LastName"

	emailParam := models.EmailParam{
		To:        "apaingha@gmail.com",
		Subject:   "test mail ",
		BodyParam: emailBody,
	}
	
	cfg.SenderEmail = "info@paingha.tech"
	sendEmail(emailParam)
}
