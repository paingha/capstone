// Copyright 2020 Paingha Joe Alagoa. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package service

import (
	"context"
	"encoding/json"

	"bitbucket.com/irb/api/config"
	"github.com/streadway/amqp"
)

//QueueService - Queue struct
type QueueService struct {
	ctx      context.Context
	cancel   context.CancelFunc
	conn     *amqp.Connection
	channel  *amqp.Channel
	incomMsg chan []byte
}

//NewQueueService - Declare a new Rabbitmq queue
func NewQueueService(ctx context.Context, cfg *config.SystemConfig) (*QueueService, error) {
	connectionString := "amqp://" + cfg.AMQPConnectionURL + cfg.RabbitMQVhost
	conn, err := amqp.Dial(connectionString)
	if err != nil {
		return nil, err
	}
	amqpChannel, err := conn.Channel()
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithCancel(ctx)
	return &QueueService{
		conn:     conn,
		channel:  amqpChannel,
		incomMsg: make(chan []byte),
		ctx:      ctx,
		cancel:   cancel,
	}, nil
}

//Run - runs queue
func (s *QueueService) Run() {

}

//Send - Sends a new work to Rabbitmq queue
func (s *QueueService) Send(channel string, payload interface{}) error {
	queue, err := s.channel.QueueDeclare(channel, true, false, false, false, nil)
	body, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	err = s.channel.Publish("", queue.Name, false, false, amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		ContentType:  "text/plain",
		Body:         body,
	})
	return nil
}
