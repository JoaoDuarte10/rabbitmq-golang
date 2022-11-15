package controller

import (
	"net/http"
	"rabbitmq-golang/src/application/services"
	"rabbitmq-golang/src/infra/logger"
)

type Controller interface {
	CreateOrder(w http.ResponseWriter, r *http.Request)
	FetchOrders(w http.ResponseWriter, r *http.Request)
}

type ControllerAdapter struct {
	Service services.OrderService
	Logger  logger.Logger
}
