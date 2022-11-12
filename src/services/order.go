package services

import "rabbitmq-golang/src/infra/http/dto"

type OrderService interface {
	CreateOrderEvent(message dto.OrderDto) error
}
