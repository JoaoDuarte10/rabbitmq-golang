package order

import (
	"fmt"
	"log"
	"net/http"
	"rabbitmq-golang/src/infra/amqp"
	"rabbitmq-golang/src/infra/http/order/controller"
	"rabbitmq-golang/src/services"
)

type OrderServer struct {
	Controller controller.Controller
	Service    services.OrderService
	http.Handler
}

func MakeOrderServer() *OrderServer {
	rabbitMQ := amqp.RabbitMQ{Uri: "amqp://example:123456@localhost:5672/"}
	service := services.OrderServiceAdapter{RabbitMQ: &rabbitMQ}
	controller := controller.ControllerAdapter{Service: &service}

	server := &OrderServer{Service: &service, Controller: &controller}

	router := &Router{controller: server.Controller}
	server.Handler = router.Init()

	return server
}

func StartServer() {
	const PORT = "3000"
	server := MakeOrderServer()

	go func() {
		err := http.ListenAndServe(fmt.Sprintf(":%s", PORT), server)
		if err != nil {
			log.Fatalf("Failed to start server. Error: %s", err)
		}
	}()

	log.Printf("Server is running on port: %s", PORT)
}
