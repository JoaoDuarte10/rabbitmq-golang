package amqp

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Worker struct{}

func (w *Worker) Start(queueName string) {
	log.Print("Starting Worker")

	rabbitMQ := RabbitMQ{}
	ch := rabbitMQ.OpenChannel()
	defer ch.Close()

	out := make(chan amqp.Delivery)

	go rabbitMQ.Consume(ch, out, queueName)

	for message := range out {
		log.Print(string(message.Body))
	}
}