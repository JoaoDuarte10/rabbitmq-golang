package factories

import (
	"rabbitmq-golang/src/application/services"
	"rabbitmq-golang/src/infra/amqp"
	"rabbitmq-golang/src/infra/amqp/workers/order"
	"rabbitmq-golang/src/infra/repository"

	"github.com/rabbitmq/amqp091-go"
)

func MakeOrderCreateWorker(qtdWorkers int) {
	db := MakeConnectionDatabse()
	repository := repository.OrderRepositorySqlite{Db: &db}
	service := services.OrderCreateService{Repository: &repository}

	channel := amqp.OpenChannel("amqp://example:123456@rabbitmq:5672/")
	rabbitMQ := amqp.RabbitMQ{
		Channel: *channel,
	}

	handler := order.HandleMessage{Service: service}

	worker := &order.OrderCreateWorker{RabbitMQ: rabbitMQ, HandleMessage: &handler}

	queueName := "create-order"
	for i := 1; i <= qtdWorkers; i++ {
		go worker.Start(queueName, 3)
	}
}

func MakeInfraRabbitMQ() {
	channel := amqp.OpenChannel("amqp://example:123456@rabbitmq:5672/")
	rabbitMQ := amqp.RabbitMQ{
		Channel: *channel,
	}

	rabbitMQ.CreateExchange(
		"order",
		"fanout",
	)

	rabbitMQ.CreateExchange(
		"order-dlx",
		"fanout",
	)

	rabbitMQ.CreateQueue(
		"create-order",
		true,
		amqp091.Table{
			"x-dead-letter-exchange":    "order",
			"x-dead-letter-routing-key": "order-create",
			"x-message-ttl":             5000,
		},
	)

	rabbitMQ.CreateQueue(
		"create-order-dlq",
		true,
		amqp091.Table{
			"x-dead-letter-exchange":    "order-dlx",
			"x-dead-letter-routing-key": "order-create",
			"x-message-ttl":             5000,
		},
	)

	rabbitMQ.QueueBind(
		"create-order",
		"order-create",
		"order-dlx",
		amqp091.Table{
			"x-dead-letter-exchange": "order-dlx",
		},
	)

	rabbitMQ.QueueBind(
		"create-order-dlq",
		"order-create",
		"order",
		amqp091.Table{
			"x-dead-letter-exchange": "order",
		},
	)
}
