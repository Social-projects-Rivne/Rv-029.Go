package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/gorilla/websocket"
	"log"
)

func InitConsumer(conn *websocket.Conn, topic string) {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	brokers := []string{"127.0.0.1:9092"}

	master, err := sarama.NewConsumer(brokers, config)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := master.Close(); err != nil {
			panic(err)
		}
	}()

	consumer, err := master.ConsumePartition(topic, 0, sarama.OffsetOldest)
	if err != nil {
		panic(err)
	}

	for {
		select {
		case err := <-consumer.Errors():
			log.Fatal(err)
		case m := <-consumer.Messages():
			err := conn.WriteMessage(1, []byte(m.Value))
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}
}
