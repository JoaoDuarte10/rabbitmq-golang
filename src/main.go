package main

import (
	"os"
	"rabbitmq-golang/src/factories"
	"rabbitmq-golang/src/infra/http/order"
)

func main() {
	forever := make(chan bool)

	factories.MakeTables()
	factories.MakeInfraRabbitMQ()

	rabbitUri := os.Getenv("RABBITMQ_BASE_URI")
	factories.MakeOrderCreateWorker(rabbitUri, 1)
	order.StartServer()

	<-forever
}
