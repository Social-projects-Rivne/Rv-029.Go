package kafka

import (
	"github.com/Shopify/sarama"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/utils/helpers"
	"github.com/gorilla/websocket"
	"net/http"
	"strconv"
	"time"
)

var Producer sarama.SyncProducer

func InitProducer() {
	config := sarama.NewConfig()
	config.Producer.Retry.Max = 10
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true

	producer, err := sarama.NewSyncProducer([]string{"127.0.0.1:9092"}, config)
	if err != nil {
		panic(err)
	}

	Producer = producer
}

func RunProducer(conn *websocket.Conn, w http.ResponseWriter) {
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			response := helpers.Response{
				Message: "Failed to store data",
				StatusCode: http.StatusInternalServerError,
			}

			response.Failed(w)
			return
		}

		message := string(msg)

		strTime := strconv.Itoa(int(time.Now().Unix()))
		producerMessage := &sarama.ProducerMessage{
			Topic: "test-topic-1",
			Key:   sarama.StringEncoder(strTime),
			Value: sarama.StringEncoder(message),
		}

		_, _, err = Producer.SendMessage(producerMessage)
		if err != nil {
			response := helpers.Response{
				Message: "Failed to store data",
				StatusCode: http.StatusInternalServerError,
			}

			response.Failed(w)
		}
	}
}
