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
	"github.com/sfreiberg/gotwilio"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

var (
	cfg config.SystemConfig
)

func handleError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

type messages struct {
	*models.Message
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
	queue, err := amqpChannel.QueueDeclare("sms", true, false, false, false, nil)
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
			var smsParam messages
			err := json.Unmarshal(d.Body, &smsParam)
			if err != nil {
				log.Printf("Error decoding JSON: %s", err)
			}
			twilio := gotwilio.NewTwilioClient(cfg.TwilioAccountSid, cfg.TwilioAuthToken)
			sendingSms(smsParam, twilio)
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

//SendWhatsapp - Sends Whatsapp text message
func (m *messages) SendWhatsapp(credentials *gotwilio.Twilio) (*gotwilio.SmsResponse, error) {
	resp, _, err := credentials.SendWhatsApp(cfg.SenderPhone, m.To, m.Content, "", "")
	if err != nil {
		return nil, err
	}
	return resp, nil
}

//SendSms - Sends Sms text message
func (m *messages) SendSms(credentials *gotwilio.Twilio) (*gotwilio.SmsResponse, error) {
	resp, _, err := credentials.SendSMS(cfg.SenderPhone, m.To, m.Content, "", "")
	if err != nil {
		return nil, err
	}
	return resp, nil
}
func sendingSms(smsParam messages, twilio *gotwilio.Twilio) {
	logrus.Infof("Sending to %s ... ", smsParam.To)
	if smsParam.Medium == "whatsapp" {
		if _, err := smsParam.SendWhatsapp(twilio); err != nil {
			logrus.Errorf("whatsapp sending error: ...", err.Error())
		}
		logrus.Info("whatsapp message sent... ")
	} else {
		if _, err := smsParam.SendSms(twilio); err != nil {
			logrus.Errorf("sms sending error: ...", err.Error())
		}
		logrus.Info("sms message sent... ")
	}

}
