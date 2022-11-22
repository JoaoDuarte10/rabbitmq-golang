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
	logger := logger.LoggerAdapter{ConsoleEnable: false}
	repository := repository.OrderRepositorySqlite{Db: &db}
	service := services.OrderCreateService{Repository: &repository, Logger: &logger}

	channel := amqp.OpenChannel(rabbitUri)
	rabbitMQ := amqp.RabbitMQ{
		Channel: *channel,
		Logger:  &logger,
	}

	handler := order.HandleMessage{Service: service, Logger: &logger}

	worker := &order.OrderCreateWorker{RabbitMQ: rabbitMQ, HandleMessage: &handler, Logger: &logger}

	queueName := "create-order"
	for i := 1; i <= qtdWorkers; i++ {
		go worker.Start(queueName, 3)
	}
}

func MakeInfraRabbitMQ() {
	logger := logger.LoggerAdapter{ConsoleEnable: false}

	rabbitUri := os.Getenv("RABBITMQ_BASE_URI")
	channel := amqp.OpenChannel(rabbitUri)
	rabbitMQ := amqp.RabbitMQ{
		Channel: *channel,
		Logger:  &logger,
	}

	exchange := "order"
	exchangeDlx := "order-dlx"
	routingKey := "order-create"
	queue := "order-create"
	queueDlq := "order-create-dlq"
	dlx := "x-dead-letter-exchange"
	dlRoutingKey := "x-dead-letter-routing-key"

	rabbitMQ.CreateExchange(
		exchange,
		"fanout",
	)

	rabbitMQ.CreateExchange(
		exchangeDlx,
		"fanout",
	)

	rabbitMQ.CreateQueue(
		queue,
		true,
		amqp091.Table{
			dlx:          exchangeDlx,
			dlRoutingKey: routingKey,
		},
	)

	rabbitMQ.CreateQueue(
		queueDlq,
		true,
		amqp091.Table{
			dlx:             exchange,
			dlRoutingKey:    routingKey,
			"x-message-ttl": 10000,
		},
	)

	rabbitMQ.QueueBind(
		queue,
		routingKey,
		exchange,
		amqp091.Table{
			dlx:          exchangeDlx,
			dlRoutingKey: routingKey,
		},
	)

	rabbitMQ.QueueBind(
		queueDlq,
		routingKey,
		exchangeDlx,
		amqp091.Table{
			dlx:          exchange,
			dlRoutingKey: routingKey,
		},
	)
}
