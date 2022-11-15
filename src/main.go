package main

import (
	"rabbitmq-golang/src/factories"
	"rabbitmq-golang/src/infra/http/order"
)

func main() {
	forever := make(chan bool)

	factories.MakeTables()
	factories.MakeInfraRabbitMQ()

	factories.MakeOrderCreateWorker(1)
	order.StartServer()

	<-forever
}
