package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/Shopify/sarama"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func KafkaSocket(w http.ResponseWriter, r *http.Request) {

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	go func() {
		for {
			select {
			case err := <-Consumer.Errors():
				log.Fatal(err)
			case m := <-Consumer.Messages():
				err = conn.WriteMessage(1, []byte(m.Value))
				if err != nil {
					fmt.Println(err)
					return
				}
			}
		}
	}()

	go func() {
		for {
			_, msg, err := conn.ReadMessage()
			if err != nil {
				fmt.Println(err)
				return
			}

			message := string(msg)

			// Initializing Kafka producer message
			strTime := strconv.Itoa(int(time.Now().Unix()))
			producerMessage := &sarama.ProducerMessage{
				Topic: "test-topic-1",
				Key:   sarama.StringEncoder(strTime),
				Value: sarama.StringEncoder(message),
			}

			// Sending message to Kafka
			_, _, _ = Producer.SendMessage(producerMessage)
		}
	}()
}
