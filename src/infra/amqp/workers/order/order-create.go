package order

import (
	"log"
	"rabbitmq-golang/src/infra/amqp"
	"rabbitmq-golang/src/infra/amqp/workers"

	"github.com/rabbitmq/amqp091-go"
)

type OrderCreateWorker struct {
	RabbitMQ amqp.RabbitMQ
	workers.HandleMessage
}

func (o *OrderCreateWorker) Start(queueName string, maxRetriesConfig int) error {
	log.Print("[OrderCreateWorker::Start] Worker Starting")

	out := make(chan amqp091.Delivery)

	go o.RabbitMQ.Consume(out, queueName)

	for message := range out {
		log.Print("[OrderCreateWorker::Consume] Message Received")

		err := o.HandleMessage.Handle(message)
		if err != nil {
			if workers.CountProcessedMessage(message) >= maxRetriesConfig {
				log.Printf(
					"[OrderCreateWorker::Consume] This order exceeded %d processing attempts",
					maxRetriesConfig,
				)
				message.Ack(true)
			} else {
				message.Nack(false, false)
			}
		} else {
			message.Ack(true)
		}
	}
	return nil
}
