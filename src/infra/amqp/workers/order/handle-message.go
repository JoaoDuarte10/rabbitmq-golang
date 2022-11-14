package order

import (
	"encoding/json"
	"rabbitmq-golang/src/application/services"
	"rabbitmq-golang/src/domain/entity"

	"github.com/rabbitmq/amqp091-go"
)

type HandleMessage struct {
	Service services.OrderCreateService
}

func (h *HandleMessage) Handle(message amqp091.Delivery) error {
	order := entity.OrderDto{}

	err := json.Unmarshal(message.Body, &order)
	if err != nil {
		return err
	}

	err = h.Service.CreateOrder(order)
	if err != nil {
		return err
	}
	return nil
}
