package workers

import (
	"encoding/json"
	"log"
	"strconv"
	"strings"

	"github.com/rabbitmq/amqp091-go"
)

type Worker interface {
	Start(queueName string)
}

func CountProcessedMessage(message amqp091.Delivery) int {
	header, err := json.Marshal(message.Headers["x-death"])
	if err != nil {
		log.Printf("[CountProcessedMessage] Error in convert header message: %s", err)
	}

	entries := strings.Split(strings.Trim(string(header), "[{}]"), ",")

	countProccess := 0

	for _, value := range entries {
		entrie := strings.Split(value, ":")
		if entrie[0] == `"count"` {
			count, _ := strconv.Atoi(entrie[1])
			countProccess = count
		}
	}
	return countProccess
}
