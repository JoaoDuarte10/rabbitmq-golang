package services

import (
	"rabbitmq-golang/src/domain/order"
	"rabbitmq-golang/src/infra/amqp"
)

type OrderServiceEvent struct {
	RabbitMQ *amqp.RabbitMQ
}

func (o *OrderServiceEvent) CreateOrderEvent(message order.OrderDto) error {
	err := o.RabbitMQ.SendMessage(message, "order-create", "order", "order-create")
	if err != nil {
		return err
	}

	return nil
}
