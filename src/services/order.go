package services

import (
	"rabbitmq-golang/src/domain/order"
)

type OrderService interface {
	CreateOrderEvent(message order.OrderDto) error
	CreateOrder(message order.OrderDto) error
}
