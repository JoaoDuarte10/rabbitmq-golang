package services

import (
	"rabbitmq-golang/src/domain/order"
	"rabbitmq-golang/src/infra/repository"
)

type GetOrderService struct {
	Repository repository.OrderRepository
}

func (g *GetOrderService) GetOrders() []order.OrderDto {
	orders, err := g.Repository.Get()
	if err != nil {
		return nil
	}
	return orders
}
