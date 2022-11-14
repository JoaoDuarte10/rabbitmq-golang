package services

import "rabbitmq-golang/src/domain/entity"

type OrderService interface {
	CreateOrderEvent(message entity.OrderDto) error
	CreateOrder(message entity.OrderDto) error
	GetOrders() []entity.OrderDto
}
