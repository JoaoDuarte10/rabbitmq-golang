package factories

import (
	"rabbitmq-golang/src/infra/amqp"
)

func MakeWorker(qtdWorkers int) {
	queueName := "golang"
	worker := amqp.Worker{}

	for i := 1; i <= qtdWorkers; i++ {
		go worker.Start(queueName)
	}
}
