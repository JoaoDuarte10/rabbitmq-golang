package order

import (
	"fmt"
	"log"
	"net/http"
	"rabbitmq-golang/src/infra/http/dto"
	"rabbitmq-golang/src/services"
)

type OrderService interface {
	Execute(message dto.OrderDto) error
}

type OrderServer struct {
	Service OrderService
	http.Handler
}

func NewOrderServer() *OrderServer {
	service := services.OrderServiceAdapter{}
	server := OrderServer{Service: &service}

	router := http.NewServeMux()
	router.Handle("/order", http.HandlerFunc(server.CreateOrder))

	server.Handler = router

	return &server
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
