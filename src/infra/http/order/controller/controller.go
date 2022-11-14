package controller

import (
	"net/http"
	"rabbitmq-golang/src/application/services"
)

type Controller interface {
	CreateOrder(w http.ResponseWriter, r *http.Request)
	FetchOrders(w http.ResponseWriter, r *http.Request)
}

type ControllerAdapter struct {
	Service services.OrderService
}
