package main

import (
	"encoding/json"
	"log"
	"os"

	"bitbucket.com/irb/api/config"
	"bitbucket.com/irb/api/models"
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

func main() {
	logrus.Info("Starting Upload Service...")
	if err := config.InitConfig(&cfg); err != nil {
		logrus.Fatalf("load config %v", err)
	}
	connectionString := "amqp://" + cfg.AMQPConnectionURL + ":5672" + cfg.RabbitMQVhost
	conn, err := amqp.Dial(connectionString)
	handleError(err, "Can't connect to AMQP")
	defer conn.Close()
	logrus.Infof("%v", cfg)
	logrus.Infof("AWS %s", cfg.AWSS3Bucket)

	amqpChannel, err := conn.Channel()
	handleError(err, "Can't create a amqpChannel")

	defer amqpChannel.Close()
	queue, err := amqpChannel.QueueDeclare("upload", true, false, false, false, nil)
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
			var fileParam models.FileParam
			err := json.Unmarshal(d.Body, &fileParam)
			if err != nil {
				log.Printf("Error decoding JSON: %s", err)
			}
			if fileParam.Medium == "local" {
				writeLocal(fileParam)
			} else {
				uploadFile(fileParam)
			}
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
