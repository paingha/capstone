// Copyright 2020 Paingha Joe Alagoa. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"os"

	"bitbucket.com/irb/api/config"
	"bitbucket.com/irb/api/models"
	"bitbucket.com/irb/api/plugins"

	
	sendgrid "github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

var (
	cfg config.SystemConfig
)

func handleError(err error, msg string) {
	if err != nil {
		plugins.LogFatal("MailService", msg, err)
	}
}

func main() {
	files := map[string]string{
		"TemplateVerifyEmail": "./templates/verify.html",
		"TemplateResetEmail":  "./templates/password-reset.html",
	}
	getFilesContents(files)
	logrus.Info("Starting service ... ")
	err := config.InitConfig(&cfg)
	if err != nil {
		handleError(err, "Can't load Env config")
	}
	connectionString := "amqp://" + cfg.AMQPConnectionURL + ":5672" + cfg.RabbitMQVhost
	conn, err := amqp.Dial(connectionString)
	handleError(err, "Can't connect to AMQP")
	defer conn.Close()
	plugins.LogInfo("MailService: Sendgrid API Key", cfg.SendgridAPIKey)

	amqpChannel, err := conn.Channel()
	handleError(err, "Can't create a amqpChannel")

	defer amqpChannel.Close()
	queue, err := amqpChannel.QueueDeclare("email", true, false, false, false, nil)
	handleError(err, "Could not declare `email` queue")

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
			plugins.LogInfo("MailService: Received a message", string(d.Body))
			var emailParam models.EmailParam
			err := json.Unmarshal(d.Body, &emailParam)
			if err != nil {
				plugins.LogError("MailService", "error unmarshalling message from rabbitmq", err)
			}
			sendEmail(emailParam)
			if err := d.Ack(false); err != nil {
				plugins.LogError("MailService", "error acknowledging message", err)
			} else {
				plugins.LogInfo("MailService: Acknowledged message", "Acknowledged message")
			}
		}
	}()

	// Stop for program termination
	<-stopChan
}

func sendEmail(emailParam models.EmailParam) {
	var bodyTemplate string
	plugins.LogInfo("MailService: Sending mail to ", emailParam.To)
	if t, ok := EmailTemplates[emailParam.Template]; ok {
		bodyTemplate = t
	} else {
		bodyTemplate = EmailTemplates["TemplateVerifyEmail"]
	}
	parsedTemplate, err := template.New("template").Parse(bodyTemplate)
	if err != nil {
		plugins.LogError("MailService", "error parsing email template", err)
	}
	var buf bytes.Buffer
	err = parsedTemplate.Execute(&buf, emailParam.BodyParam)

	from := mail.NewEmail("Northeastern State University IRB", cfg.SenderEmail)
	subject := emailParam.Subject
	to := mail.NewEmail(fmt.Sprintf("%s %s ", emailParam.BodyParam["first_name"], emailParam.BodyParam["last_name"]), emailParam.To)

	htmlContent := buf.String()
	message := mail.NewSingleEmail(from, subject, to, "text", htmlContent)

	client := sendgrid.NewSendClient(cfg.SendgridAPIKey)
	response, err := client.Send(message)
	fmt.Println(response.StatusCode)
	if err != nil {
		plugins.LogError("MailService", "error sending email", err)
	} else {
		plugins.LogInfo("MailService: Email sent successfully", "200")
	}
}
