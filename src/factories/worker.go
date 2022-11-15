package factories

import (
	"os"
	"rabbitmq-golang/src/application/services"
	"rabbitmq-golang/src/infra/amqp"
	"rabbitmq-golang/src/infra/amqp/workers/order"
	"rabbitmq-golang/src/infra/logger"
	"rabbitmq-golang/src/infra/repository"

	"github.com/rabbitmq/amqp091-go"
)

func MakeOrderCreateWorker(rabbitUri string, qtdWorkers int) {
	db := MakeConnectionDatabse()
	repository := repository.OrderRepositorySqlite{Db: &db}
	service := services.OrderCreateService{Repository: &repository}

	channel := amqp.OpenChannel(rabbitUri)
	rabbitMQ := amqp.RabbitMQ{
		Channel: *channel,
	}

	logger := &logger.LoggerAdapter{ConsoleEnable: false}

	handler := order.HandleMessage{Service: service, Logger: logger}

	worker := &order.OrderCreateWorker{RabbitMQ: rabbitMQ, HandleMessage: &handler, Logger: logger}

	queueName := "create-order"
	for i := 1; i <= qtdWorkers; i++ {
		go worker.Start(queueName, 3)
	}
}

func MakeInfraRabbitMQ() {
	rabbitUri := os.Getenv("RABBITMQ_BASE_URI")
	channel := amqp.OpenChannel(rabbitUri)
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
