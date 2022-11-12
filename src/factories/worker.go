package factories

import (
	"rabbitmq-golang/src/infra/amqp"
	"rabbitmq-golang/src/infra/amqp/workers"
	"rabbitmq-golang/src/services"
)

func MakeOrderCreateWorker(qtdWorkers int) {
	queueName := "golang"
	service := services.OrderCreateService{}
	rabbitMQ := amqp.RabbitMQ{
		Uri: "amqp://example:123456@localhost:5672/",
	}
	worker := &workers.OrderCreateWorker{RabbitMQ: rabbitMQ, Service: service}

	for i := 1; i <= qtdWorkers; i++ {
		go worker.Start(queueName)
	}
}
