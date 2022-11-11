package factories

import (
	"rabbitmq-golang/src/infra/amqp"
)

func MakeWorker() {
	queueName := "golang"
	worker := amqp.Worker{}

	forever := make(chan bool)

	qtdWorkers := 5
	for i := 1; i <= qtdWorkers; i++ {
		go worker.Start(queueName)
	}
	<-forever
}
