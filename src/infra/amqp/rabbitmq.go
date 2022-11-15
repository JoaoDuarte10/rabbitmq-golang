package amqp

import (
	"context"
	"encoding/json"
	"fmt"
	"rabbitmq-golang/src/infra/logger"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQ struct {
	Channel amqp.Channel
	Logger  logger.Logger
}

func failOnError(err error, msg string) {
	logger := &logger.LoggerAdapter{ConsoleEnable: false}

	if err != nil {
		logger.Error(fmt.Sprintf("%s: %s", msg, err))
		return
	}
}

func OpenChannel(uri string) *amqp.Channel {
	conn, err := amqp.Dial(uri)
	failOnError(err, "[RabbitMQ::OpenChannel] Failed to connect to RabbitMQ")

	ch, err := conn.Channel()
	failOnError(err, "[RabbitMQ::OpenChannel] Failed to open a channel")

	return ch
}

func (r *RabbitMQ) Consume(out chan amqp.Delivery, queueName string) {
	messages, err := r.Channel.Consume(
		queueName,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "[RabbitMQ::Consume] Failed to consume messages")

	for message := range messages {
		out <- message
	}
}

func (r *RabbitMQ) SendMessage(message any, queueName string, exchange string, key string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	body, err := json.Marshal(message)
	failOnError(err, fmt.Sprintf("[RabbitMQ::SendMessage] Failed to parsed message in json: %s", message))

	err = r.Channel.PublishWithContext(ctx,
		exchange,
		key,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		})
	if err != nil {
		failOnError(err, "[RabbitMQ::SendMessage] Failed to publish a message")
		return err
	}

	r.Logger.Info(fmt.Sprintf("[RabbitMQ::SendMessage] Success publish message in queue: %s", queueName))
	return nil
}

func (r *RabbitMQ) CreateExchange(
	name string,
	exchangeType string,
) error {
	err := r.Channel.ExchangeDeclare(
		name,
		exchangeType,
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "[RabbitMQ::CreateExchange] Failed to declare exchange")
	return nil
}

func (r *RabbitMQ) CreateQueue(
	name string,
	durable bool,
	args amqp.Table,
) error {
	_, err := r.Channel.QueueDeclare(
		name,
		durable,
		false,
		false,
		false,
		args,
	)
	failOnError(err, "[RabbitMQ::CreateQueue] Failed to declare exchange")
	return nil
}

func (r *RabbitMQ) QueueBind(
	name string,
	key string,
	exchange string,
	args amqp.Table,
) error {
	err := r.Channel.QueueBind(
		name,
		key,
		exchange,
		false,
		args,
	)
	failOnError(err, "[RabbitMQ::QueueBind] Failed to declare exchange")
	return nil
}
