package repositories

import amqp "github.com/rabbitmq/amqp091-go"

type ISuscriber interface {
	ReceiveContent(nameQueue, routingKey string) <-chan amqp.Delivery
}