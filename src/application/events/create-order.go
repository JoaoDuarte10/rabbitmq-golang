package events

import (
	"rabbitmq-golang/src/domain/entity"
	"rabbitmq-golang/src/infra/amqp"
)

type OrderServiceEvent struct {
	RabbitMQ *amqp.RabbitMQ
}

func (o *OrderServiceEvent) CreateOrderEvent(message entity.OrderDto) error {
	err := o.RabbitMQ.SendMessage(message, "create-order", "order", "order-create")
	if err != nil {
		return err
	}

	return nil
}
