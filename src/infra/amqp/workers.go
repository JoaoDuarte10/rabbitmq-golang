package amqp

type Worker interface {
	Start(queueName string)
}
