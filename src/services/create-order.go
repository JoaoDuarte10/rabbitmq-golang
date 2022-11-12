package services

import (
	"log"
	dto "rabbitmq-golang/src/domain/order"
	"rabbitmq-golang/src/infra/repository"

	_ "github.com/mattn/go-sqlite3"
)

type OrderCreateService struct {
	Repository repository.OrderRepository
}

func (o *OrderCreateService) CreateOrder(order dto.OrderDto) error {
	err := o.Repository.Save(order)
	if err != nil {
		log.Printf("Error in save order: %s", err)
		return err
	}
	return nil
}
