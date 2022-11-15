package controller

import (
	"encoding/json"
	"io"
	"net/http"
	"rabbitmq-golang/src/domain/entity"
)

func (c *ControllerAdapter) CreateOrder(w http.ResponseWriter, r *http.Request) {
	order := entity.OrderDto{}

	body, _ := io.ReadAll(r.Body)

	err := json.Unmarshal(body, &order)
	if err != nil {
		c.Logger.Error("[ControllerAdapter::CreateOrder] Failed to convert body in json")
		w.WriteHeader(400)
	}

	err = c.Service.CreateOrderEvent(order)
	if err != nil {
		c.Logger.Error("[ControllerAdapter::CreateOrder] Failed to dispath order event")
		w.WriteHeader(500)
	}

	w.WriteHeader(201)
}
