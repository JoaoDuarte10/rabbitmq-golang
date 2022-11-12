package order

import (
	"net/http"
	"rabbitmq-golang/src/infra/http/order/controller"
)

type Router struct {
	controller controller.Controller
	http.Handler
}

func (r *Router) Init() *http.ServeMux {
	router := http.NewServeMux()
	router.Handle("/order", http.HandlerFunc(r.controller.CreateOrder))
	return router
}
