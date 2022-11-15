package order

import (
	"fmt"
	"rabbitmq-golang/src/infra/amqp"
	"rabbitmq-golang/src/infra/amqp/workers"
	"rabbitmq-golang/src/infra/logger"

	"github.com/rabbitmq/amqp091-go"
)

type OrderCreateWorker struct {
	RabbitMQ amqp.RabbitMQ
	workers.HandleMessage
	Logger logger.Logger
}

func (o *OrderCreateWorker) Start(queueName string, maxRetriesConfig int) error {
	o.Logger.Info("[OrderCreateWorker::Start] Worker Starting")

	out := make(chan amqp091.Delivery)

	go o.RabbitMQ.Consume(out, queueName)

	for message := range out {
		o.Logger.Info("[OrderCreateWorker::Consume] Message Received")

		err := o.HandleMessage.Handle(message)
		if err != nil {
			if workers.CountProcessedMessage(message) >= maxRetriesConfig {
				o.Logger.Error(fmt.Sprintf(
					"[OrderCreateWorker::Consume] This order exceeded %d processing attempts",
					maxRetriesConfig,
				))
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
