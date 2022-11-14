package workers

import "github.com/rabbitmq/amqp091-go"

type HandleMessage interface {
	Handle(message amqp091.Delivery) error
}
