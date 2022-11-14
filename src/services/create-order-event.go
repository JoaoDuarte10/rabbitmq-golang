package services

import (
	"rabbitmq-golang/src/domain/order"
	"rabbitmq-golang/src/infra/amqp"
)

type OrderServiceEvent struct {
	RabbitMQ *amqp.RabbitMQ
}

func (o *OrderServiceEvent) CreateOrderEvent(message order.OrderDto) error {
	channel := o.RabbitMQ.OpenChannel()
	defer channel.Close()

	err := o.RabbitMQ.SendMessage(channel, message, "order-create", "order", "order-create")
	if err != nil {
		return err
	}

	return nil
}
