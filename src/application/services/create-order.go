package services

import (
	"fmt"
	"rabbitmq-golang/src/domain/entity"
	"rabbitmq-golang/src/infra/logger"
	"rabbitmq-golang/src/infra/repository"

	_ "github.com/mattn/go-sqlite3"
)

type OrderCreateService struct {
	Repository repository.OrderRepository
	Logger     logger.Logger
}

func (o *OrderCreateService) CreateOrder(order entity.OrderDto) error {
	err := o.Repository.Save(order)
	if err != nil {
		o.Logger.Error(fmt.Sprintf("[OrderCreateService::CreateOrder] Error in save order: %s", err))
		return err
	}
	return nil
}
