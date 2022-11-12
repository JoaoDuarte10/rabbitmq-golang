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
	router.Handle("/order", http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if r.ValidateMethod(req, http.MethodPost) {
			r.controller.CreateOrder(w, req)
		} else {
			w.WriteHeader(404)
		}
	}))
	return router
}

func (r *Router) ValidateMethod(req *http.Request, method string) bool {
	return req.Method == method
}
