package factories

import (
	"rabbitmq-golang/src/infra/amqp"
	"rabbitmq-golang/src/infra/amqp/workers/order"
	"rabbitmq-golang/src/infra/repository"
	"rabbitmq-golang/src/services"
)

func MakeOrderCreateWorker(qtdWorkers int) {
	db := MakeConnectionDatabse()
	repository := repository.OrderRepositorySqlite{Db: &db}
	service := services.OrderCreateService{Repository: &repository}

	queueName := "order-create"
	rabbitMQ := amqp.RabbitMQ{
		Uri: "amqp://example:123456@localhost:5672/",
	}

	handler := order.HandleMessage{Service: service}

	worker := &order.OrderCreateWorker{RabbitMQ: rabbitMQ, HandleMessage: &handler}

	for i := 1; i <= qtdWorkers; i++ {
		go worker.Start(queueName, 3)
	}
}
