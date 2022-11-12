package order

import (
	"fmt"
	"log"
	"net/http"
	"rabbitmq-golang/src/infra/amqp"
	"rabbitmq-golang/src/infra/http/order/controller"
	"rabbitmq-golang/src/infra/repository"
	"rabbitmq-golang/src/services"
)

type OrderServer struct {
	Controller controller.Controller
	Service    services.OrderService
	http.Handler
}

type OrderServiceAdapter struct {
	*amqp.RabbitMQ
	*services.OrderCreateService
	*services.OrderServiceEvent
	*services.GetOrderService
}

func MakeOrderServer() *OrderServer {
	rabbitMQ := amqp.RabbitMQ{Uri: "amqp://example:123456@localhost:5672/"}
	orderCreateEvent := services.OrderServiceEvent{RabbitMQ: &rabbitMQ}
	orderCreateService := services.OrderCreateService{Repository: &repository.OrderRepositorySqlite{}}
	fetchOrders := services.GetOrderService{Repository: &repository.OrderRepositorySqlite{}}

	service := OrderServiceAdapter{
		RabbitMQ:           &rabbitMQ,
		OrderCreateService: &orderCreateService,
		OrderServiceEvent:  &orderCreateEvent,
		GetOrderService:    &fetchOrders,
	}

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
