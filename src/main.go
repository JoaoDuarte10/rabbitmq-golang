package main

import (
	"rabbitmq-golang/src/infra/amqp"
)

func main() {
	queueName := "golang"
	worker := amqp.Worker{}
	worker.Start(queueName)
}
