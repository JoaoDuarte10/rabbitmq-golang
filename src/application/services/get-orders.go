package services

import (
	"rabbitmq-golang/src/domain/entity"
	"rabbitmq-golang/src/infra/repository"
)

type GetOrderService struct {
	Repository repository.OrderRepository
}

func (g *GetOrderService) GetOrders() []entity.OrderDto {
	orders, err := g.Repository.Get()
	if err != nil {
		return nil
	}
	return orders
}
