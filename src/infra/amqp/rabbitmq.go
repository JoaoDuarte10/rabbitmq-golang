package amqp

import (
	"context"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

type RabbitMQ struct{}

const AMQP_URI = "amqp://example:123456@localhost:5672/"

func (r *RabbitMQ) OpenChannel() *amqp.Channel {
	conn, err := amqp.Dial(AMQP_URI)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	return ch
}

func (r *RabbitMQ) Consume(ch *amqp.Channel, out chan amqp.Delivery, queueName string) {
	_, err := ch.QueueDeclare(queueName, false, false, false, false, nil)
	failOnError(err, "Failed to declare a queue")

	messages, err := ch.Consume(
		queueName,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to declare a queue")

	for message := range messages {
		out <- message
	}
}

func (r *RabbitMQ) sendMessage(ch *amqp.Channel, message string, queueName string, exchange string) {
	q, err := ch.QueueDeclare(queueName, false, false, false, false, nil)
	failOnError(err, "Failed to declare a queue")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = ch.PublishWithContext(ctx,
		exchange,
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
	failOnError(err, "Failed to publish a message")
	log.Printf("Success publish message in queue: %s", queueName)
}
