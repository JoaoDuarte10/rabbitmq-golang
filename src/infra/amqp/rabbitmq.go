package amqp

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

type RabbitMQ struct {
	Uri string
}

func (r *RabbitMQ) OpenChannel() *amqp.Channel {
	conn, err := amqp.Dial(r.Uri)
	failOnError(err, "[RabbitMQ::OpenChannel] Failed to connect to RabbitMQ")

	ch, err := conn.Channel()
	failOnError(err, "[RabbitMQ::OpenChannel] Failed to open a channel")

	return ch
}

func (r *RabbitMQ) Consume(ch *amqp.Channel, out chan amqp.Delivery, queueName string) {
	_, err := ch.QueueDeclare(queueName, false, false, false, false, nil)
	failOnError(err, "[RabbitMQ::Consume] Failed to declare a queue")

	messages, err := ch.Consume(
		queueName,
		"",
		true,
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

func (r *RabbitMQ) SendMessage(ch *amqp.Channel, message string, queueName string, exchange string) error {
	q, err := ch.QueueDeclare(queueName, false, false, false, false, nil)
	failOnError(err, "Failed to declare a queue")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	body, err := json.Marshal(message)
	failOnError(err, fmt.Sprintf("[RabbitMQ::SendMessage] Failed to parsed message in json: %s", message))

	err = ch.PublishWithContext(ctx,
		exchange,
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		})
	if err != nil {
		failOnError(err, "Failed to publish a message")
		return err
	}

	log.Printf("Success publish message in queue: %s", queueName)
	return nil
}
