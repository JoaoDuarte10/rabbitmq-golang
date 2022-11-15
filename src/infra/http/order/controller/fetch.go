package controller

import (
	"encoding/json"
	"net/http"
)

func (c *ControllerAdapter) FetchOrders(w http.ResponseWriter, r *http.Request) {
	orders := c.Service.GetOrders()

	err := json.NewEncoder(w).Encode(orders)
	if err != nil {
		w.WriteHeader(500)
	}
}
