package workers

import (
	"encoding/json"
	"log"
	"rabbitmq-golang/src/domain/order"
	"rabbitmq-golang/src/infra/amqp"
	"rabbitmq-golang/src/services"

	"github.com/rabbitmq/amqp091-go"
)

type OrderCreateWorker struct {
	RabbitMQ amqp.RabbitMQ
	Service  services.OrderCreateService
}

func (o *OrderCreateWorker) Start(queueName string) error {
	log.Print("Starting OrderCreateWorker...")

	ch := o.RabbitMQ.OpenChannel()
	defer ch.Close()

	out := make(chan amqp091.Delivery)

	go o.RabbitMQ.Consume(ch, out, queueName)

	for message := range out {
		order := order.OrderDto{}

		err := json.Unmarshal(message.Body, &order)
		if err != nil {
			log.Fatal("Deu ruim no worker")
		}

		err = o.Service.CreateOrder(order)
		if err != nil {
			log.Fatal("Failed to process message")
			message.Ack(false)
			return err
		}
		message.Ack(true)
	}
	return nil
}
