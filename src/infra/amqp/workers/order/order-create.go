package order

import (
	"fmt"
	"log"
	"rabbitmq-golang/src/infra/amqp"
	"rabbitmq-golang/src/infra/amqp/workers"
	"strconv"
	"strings"

	"github.com/rabbitmq/amqp091-go"
)

type OrderCreateWorker struct {
	RabbitMQ amqp.RabbitMQ
	workers.HandleMessage
}

func (o *OrderCreateWorker) Start(queueName string, maxRetriesConfig int) error {
	log.Print("[OrderCreateWorker::Start] Worker Starting")

	ch := o.RabbitMQ.OpenChannel()
	defer ch.Close()

	out := make(chan amqp091.Delivery)

	go o.RabbitMQ.Consume(ch, out, queueName)

	for message := range out {
		log.Print("[OrderCreateWorker::Consume] Process Message...")
		err := o.HandleMessage.Handle(message)
		if err != nil {
			if CountProcessedMessage(message) >= maxRetriesConfig {
				log.Printf("[OrderCreateWorker::Consume] This order exceeded %d processing attempts", maxRetriesConfig)
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

func CountProcessedMessage(message amqp091.Delivery) int {
	header := message.Headers["x-death"]

	result := (fmt.Sprint(header))

	entries := strings.Split(result, "map[count:")

	countProccess := 0

	for key, value := range entries {
		if key == 1 {
			count, err := strconv.Atoi(strings.Split(value, " ")[0])
			if err != nil {
				log.Printf("[CountProcessedMessage] Error in get message counter processing: %s", err)
			}
			countProccess = count
		}
	}
	return countProccess
}
