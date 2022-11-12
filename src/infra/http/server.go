package http

import (
	"fmt"
	"log"
	"net/http"

	"rabbitmq-golang/src/infra/http/controller"
)

func StartServer() {
	const PORT = "3000"

	http.HandleFunc("/send", func(w http.ResponseWriter, r *http.Request) {
		controller.SendMessage(w, r)
	})

	go func() {
		err := http.ListenAndServe(fmt.Sprintf(":%s", PORT), nil)
		if err != nil {
			log.Fatalf("Failed to start server. Error: %s", err)
		}
	}()

	log.Printf("Server is running on port: %s", PORT)
}
