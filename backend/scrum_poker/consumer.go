package scrum_poker

import (
	"github.com/Shopify/sarama"
	"log"
	"encoding/json"
	"github.com/gocql/gocql"
	"fmt"
)

var Consumer sarama.PartitionConsumer

func InitConsumer(topic string, partition int32) {

	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	brokers := []string{"127.0.0.1:9092"}

	master, err := sarama.NewConsumer(brokers, config)
	if err != nil {
		panic(err)
	}

	consumer, err := master.ConsumePartition(topic, partition, sarama.OffsetOldest)
	if err != nil {
		panic(err)
	}

	go func() {
		select {

		case err := <-consumer.Errors():
			log.Fatal(err)

		case m := <-consumer.Messages():
			message := []byte(m.Value)

			fmt.Println(string(message))
			jsonMessage := make(map[string]interface{}, 0)
			err := json.Unmarshal(message, jsonMessage)
			if err != nil {
				//TODO: error
			}

			if sprintID, ok := jsonMessage["sprintID"]; ok {
				sprintUUI, err := gocql.ParseUUID(sprintID.(string))
				if err != nil {
					//TODO: error
				}

				if hub, ok := ActiveHubs[sprintUUI]; ok {
					hub.Broadcast <- message
				}
			}
		}
	}()

	Consumer = consumer
}

