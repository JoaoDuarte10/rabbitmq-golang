package services

import (
	"log"
	"rabbitmq-golang/src/domain/order"
)

type OrderCreateService struct{}

func (o *OrderCreateService) CreateOrder(order order.OrderDto) error {
	log.Print("Service")
	return nil
}
