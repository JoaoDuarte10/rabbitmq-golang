package main

import (
	"rabbitmq-golang/src/factories"
	"rabbitmq-golang/src/infra/http"
)

func main() {
	forever := make(chan bool)

	factories.MakeWorker()
	http.StartServer()

	<-forever
}
