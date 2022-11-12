package order

import (
	"fmt"
	"log"
	"net/http"
	"rabbitmq-golang/src/infra/amqp"
)

type OrderService interface {
	Execute(message OrderDto) error
}

type OrderServer struct {
	Service OrderService
	http.Handler
}

func (o *OrderServer) Execute(message OrderDto) error {
	rabbitMQ := amqp.RabbitMQ{
		Uri: "amqp://example:123456@localhost:5672/",
	}
	ch := rabbitMQ.OpenChannel()
	defer ch.Close()

	err := rabbitMQ.SendMessage(ch, message, "golang", "")
	if err != nil {
		return err
	}

	return nil
}

func NewOrderServer() *OrderServer {
	server := new(OrderServer)

	router := http.NewServeMux()
	router.Handle("/order", http.HandlerFunc(server.CreateOrder))

	server.Handler = router

	return server
}

func StartServer() {
	const PORT = "3000"
	server := NewOrderServer()

	go func() {
		err := http.ListenAndServe(fmt.Sprintf(":%s", PORT), server)
		if err != nil {
			log.Fatalf("Failed to start server. Error: %s", err)
		}
	}()

	log.Printf("Server is running on port: %s", PORT)
}
