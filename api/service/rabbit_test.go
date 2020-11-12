// Copyright 2020 Paingha Joe Alagoa. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package service

import (
	"context"
	"testing"

	"bitbucket.com/irb/api/config"
	"bitbucket.com/irb/api/models"
	"github.com/stretchr/testify/assert"
)

func TestQueueService_Send(t *testing.T) {

	cfg := config.SystemConfig{
		AMQPConnectionURL: "amqp://localhost/",
	}
	service, err := NewQueueService(context.Background(), &cfg)
	assert.Nil(t, err)
	payload := models.EmailParam{
		To:      "apaingha@gmail.com",
		Subject: "Verify your email",
		BodyParam: map[string]string{
			"first_name": "Joe",
			"last_name":  "Alagoa",
		},
	}
	err = service.Send("email", payload)
	assert.Nil(t, err)
}
