package controller

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"rabbitmq-golang/src/infra/http/dto"
)

func (c *ControllerAdapter) CreateOrder(w http.ResponseWriter, r *http.Request) {
	order := dto.OrderDto{}

	body, _ := io.ReadAll(r.Body)

	err := json.Unmarshal(body, &order)
	if err != nil {
		log.Print("Failed to convert body in json")
		w.WriteHeader(400)
	}

	err = c.Service.CreateOrderEvent(order)
	if err != nil {
		log.Print("Erro no service")
	}

	w.WriteHeader(201)
}
