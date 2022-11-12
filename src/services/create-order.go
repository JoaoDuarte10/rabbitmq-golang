package services

import (
	"rabbitmq-golang/src/infra/amqp"
	"rabbitmq-golang/src/infra/http/dto"
)

type OrderServiceAdapter struct {
	RabbitMQ *amqp.RabbitMQ
}

func (o *OrderServiceAdapter) CreateOrderEvent(message dto.OrderDto) error {
	channel := o.RabbitMQ.OpenChannel()
	defer channel.Close()

	err := o.RabbitMQ.SendMessage(channel, message, "golang", "")
	if err != nil {
		return err
	}

	return nil
}
