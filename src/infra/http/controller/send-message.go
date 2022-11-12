package controller

import (
	"io/ioutil"
	"log"
	"net/http"

	amqp "rabbitmq-golang/src/infra/amqp"
)

func SendMessage(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal("[Http::SendMessage] Failed parsed body in json")
	}

	rabbitMQ := amqp.RabbitMQ{
		Uri: "amqp://example:123456@localhost:5672/",
	}
	ch := rabbitMQ.OpenChannel()
	defer ch.Close()
	err = rabbitMQ.SendMessage(ch, string(body), "golang", "")

	if err != nil {
		w.WriteHeader(500)
	}

	w.WriteHeader(200)
}
