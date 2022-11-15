package order

import (
	"fmt"
	"log"
	"net/http"
	"rabbitmq-golang/src/application/events"
	"rabbitmq-golang/src/application/services"
	"rabbitmq-golang/src/factories"
	"rabbitmq-golang/src/infra/amqp"
	"rabbitmq-golang/src/infra/http/order/controller"
	"rabbitmq-golang/src/infra/repository"
)

type OrderServer struct {
	Controller controller.Controller
	Service    services.OrderService
	http.Handler
}

type OrderServiceAdapter struct {
	*amqp.RabbitMQ
	*services.OrderCreateService
	*events.OrderServiceEvent
	*services.GetOrderService
}

func MakeOrderServer() *OrderServer {
	channel := amqp.OpenChannel("amqp://example:123456@rabbitmq:5672/")
	rabbitMQ := amqp.RabbitMQ{Channel: *channel}
	db := factories.MakeConnectionDatabse()
	repository := &repository.OrderRepositorySqlite{Db: &db}

	orderCreateEvent := events.OrderServiceEvent{RabbitMQ: &rabbitMQ}
	orderCreateService := services.OrderCreateService{Repository: repository}
	fetchOrders := services.GetOrderService{Repository: repository}

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

	log.Printf("[OrderServer::StartServer] Server is running on port: %s", PORT)
}
