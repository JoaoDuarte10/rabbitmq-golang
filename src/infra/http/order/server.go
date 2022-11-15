package order

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"rabbitmq-golang/src/application/events"
	"rabbitmq-golang/src/application/services"
	"rabbitmq-golang/src/factories"
	"rabbitmq-golang/src/infra/amqp"
	"rabbitmq-golang/src/infra/http/order/controller"
	"rabbitmq-golang/src/infra/logger"
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

func MakeOrderServer(rabbitUri string) *OrderServer {
	channel := amqp.OpenChannel(rabbitUri)
	rabbitMQ := amqp.RabbitMQ{Channel: *channel}
	db := factories.MakeConnectionDatabse()
	repository := &repository.OrderRepositorySqlite{Db: &db}

	logger := &logger.LoggerAdapter{ConsoleEnable: false}

	orderCreateEvent := events.OrderServiceEvent{RabbitMQ: &rabbitMQ}
	orderCreateService := services.OrderCreateService{Repository: repository, Logger: logger}
	fetchOrders := services.GetOrderService{Repository: repository}

	service := OrderServiceAdapter{
		RabbitMQ:           &rabbitMQ,
		OrderCreateService: &orderCreateService,
		OrderServiceEvent:  &orderCreateEvent,
		GetOrderService:    &fetchOrders,
	}

	controller := controller.ControllerAdapter{Service: &service, Logger: logger}

	server := &OrderServer{Service: &service, Controller: &controller}

	router := &Router{controller: server.Controller}
	server.Handler = router.Init()

	return server
}

func StartServer() {
	const PORT = "3000"
	rabbitUri := os.Getenv("RABBITMQ_BASE_URI")
	server := MakeOrderServer(rabbitUri)

	go func() {
		err := http.ListenAndServe(fmt.Sprintf(":%s", PORT), server)
		if err != nil {
			log.Fatalf("[StartServer] Failed to start server. Error: %s", err)
		}
	}()

	logger := &logger.LoggerAdapter{ConsoleEnable: false}

	logger.Info(fmt.Sprintf("[OrderServer::StartServer] Server is running on port: %s", PORT))
}
