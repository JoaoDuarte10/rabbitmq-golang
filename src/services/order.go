package services

import "rabbitmq-golang/src/infra/http/dto"

type OrderService interface {
	CreateOrder(message dto.OrderDto) error
}
