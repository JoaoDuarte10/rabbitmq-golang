package main

import (
	"rabbitmq-golang/src/factories"
	"rabbitmq-golang/src/infra/http/order"
)

func main() {
	forever := make(chan bool)

	factories.MakeOrderCreateWorker(2)
	order.StartServer()

	<-forever
}
