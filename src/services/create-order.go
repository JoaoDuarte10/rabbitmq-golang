package services

import (
	"rabbitmq-golang/src/infra/amqp"
	"rabbitmq-golang/src/infra/http/dto"
)

type OrderServiceAdapter struct{}

func (o *OrderServiceAdapter) Execute(message dto.OrderDto) error {
	rabbitMQ := amqp.RabbitMQ{
		Uri: "amqp://example:123456@localhost:5672/",
	}
	channel := rabbitMQ.OpenChannel()
	defer channel.Close()

	err := rabbitMQ.SendMessage(channel, message, "golang", "")
	if err != nil {
		return err
	}

	return nil
}
