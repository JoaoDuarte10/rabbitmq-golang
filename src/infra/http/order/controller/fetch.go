package controller

import (
	"encoding/json"
	"net/http"
)

func (c *ControllerAdapter) FetchOrders(w http.ResponseWriter, r *http.Request) {
	orders := c.Service.GetOrders()

	json.NewEncoder(w).Encode(orders)
	w.WriteHeader(http.StatusOK)
}
