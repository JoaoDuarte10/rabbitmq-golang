package factories

import (
	"rabbitmq-golang/src/application/services"
	"rabbitmq-golang/src/infra/amqp"
	"rabbitmq-golang/src/infra/amqp/workers/order"
	"rabbitmq-golang/src/infra/repository"
)

func MakeOrderCreateWorker(qtdWorkers int) {
	db := MakeConnectionDatabse()
	repository := repository.OrderRepositorySqlite{Db: &db}
	service := services.OrderCreateService{Repository: &repository}

	channel := amqp.OpenChannel("amqp://example:123456@localhost:5672/")
	queueName := "order-create"
	rabbitMQ := amqp.RabbitMQ{
		Channel: *channel,
	}

	handler := order.HandleMessage{Service: service}

	worker := &order.OrderCreateWorker{RabbitMQ: rabbitMQ, HandleMessage: &handler}

	for i := 1; i <= qtdWorkers; i++ {
		go worker.Start(queueName, 3)
	}
}
