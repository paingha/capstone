// Copyright 2020 Paingha Joe Alagoa. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"encoding/json"
	"log"
	"os"

	"bitbucket.com/irb/api/config"
	"bitbucket.com/irb/api/models"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"github.com/tbalthazar/onesignal-go"
)

var (
	cfg config.SystemConfig
)

func handleError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

type push struct {
	*models.Push
}

func main() {
	logrus.Info("Starting service ... ")
	err := config.InitConfig(&cfg)
	if err != nil {
		logrus.Fatalf("load config %v", err)
	}
	conn, err := amqp.Dial(cfg.AMQPConnectionURL)
	handleError(err, "Can't connect to AMQP")
	defer conn.Close()
	logrus.Infof("%v", cfg)
	logrus.Infof("Sender number %s", cfg.SenderPhone)

	amqpChannel, err := conn.Channel()
	handleError(err, "Can't create a amqpChannel")

	defer amqpChannel.Close()
	queue, err := amqpChannel.QueueDeclare("push", true, false, false, false, nil)
	handleError(err, "Could not declare `add` queue")

	err = amqpChannel.Qos(1, 0, false)
	handleError(err, "Could not configure QoS")

	messageChannel, err := amqpChannel.Consume(
		queue.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	handleError(err, "Could not register consumer")

	stopChan := make(chan bool)

	go func() {
		log.Printf("Consumer ready, PID: %d", os.Getpid())
		for d := range messageChannel {
			log.Printf("Received a message: %s", d.Body)
			var pushParam push
			err := json.Unmarshal(d.Body, &pushParam)
			if err != nil {
				log.Printf("Error decoding JSON: %s", err)
			}
			client := onesignal.NewClient(nil)
			client.AppKey = cfg.OneSignalAppKey
			sendingPush(pushParam, client)
			if err := d.Ack(false); err != nil {
				log.Printf("Error acknowledging message : %s", err)
			} else {
				log.Printf("Acknowledged message")
			}
		}
	}()

	// Stop for program termination
	<-stopChan
}

func sendingPush(pushParam push, client *onesignal.Client) {
	logrus.Infof("Push content %s ... ", pushParam.Content)
	notificationReq := &onesignal.NotificationRequest{
		AppID:            cfg.AppID,
		Contents:         pushParam.Content,
		IsIOS:            pushParam.IsIOS,
		IncludePlayerIDs: pushParam.Players,
	}
	_, _, err := client.Notifications.Create(notificationReq)
	logrus.Infof("err %s ... ", err)
}
