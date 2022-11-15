package main

import (
	"os"
	"rabbitmq-golang/src/factories"
	"rabbitmq-golang/src/infra/http/order"
	"rabbitmq-golang/src/infra/logger"
)

func main() {
	forever := make(chan bool)

	logger := logger.LoggerAdapter{ConsoleEnable: false}
	logger.Info("Initialize Application")
	factories.MakeTables()
	factories.MakeInfraRabbitMQ()

	rabbitUri := os.Getenv("RABBITMQ_BASE_URI")
	factories.MakeOrderCreateWorker(rabbitUri, 1)
	order.StartServer()

	<-forever
}
