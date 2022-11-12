package order

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func (o *OrderServer) CreateOrder(w http.ResponseWriter, r *http.Request) {
	order := OrderDto{}

	body, _ := io.ReadAll(r.Body)

	err := json.Unmarshal(body, &order)
	if err != nil {
		log.Print("Failed to convert body in json")
		w.WriteHeader(400)
	}

	err = o.Execute(order)
	if err != nil {
		log.Print("Erro no service")
	}

	w.WriteHeader(201)
}
