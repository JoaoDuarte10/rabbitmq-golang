package services

import (
	"log"
	"rabbitmq-golang/src/domain/entity"
	"rabbitmq-golang/src/infra/repository"

	_ "github.com/mattn/go-sqlite3"
)

type OrderCreateService struct {
	Repository repository.OrderRepository
}

func (o *OrderCreateService) CreateOrder(order entity.OrderDto) error {
	err := o.Repository.Save(order)
	if err != nil {
		log.Printf("[OrderCreateService::CreateOrder] Error in save order: %s", err)
		return err
	}
	return nil
}
