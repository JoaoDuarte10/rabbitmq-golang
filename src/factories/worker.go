package factories

import (
	"rabbitmq-golang/src/infra/amqp"
)

func MakeWorker() {
	queueName := "golang"
	worker := amqp.Worker{}

	qtdWorkers := 5
	for i := 1; i <= qtdWorkers; i++ {
		go worker.Start(queueName)
	}
}
