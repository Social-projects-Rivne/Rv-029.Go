package scrum_poker

import (
	"log"

	"github.com/Shopify/sarama"
	"github.com/gocql/gocql"

	"encoding/json"
)

var Consumer sarama.PartitionConsumer

func InitConsumer(topic string) {

	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	brokers := []string{"127.0.0.1:9092"}

	master, err := sarama.NewConsumer(brokers, config)
	if err != nil {
		panic(err)
	}

	consumer, err := master.ConsumePartition(topic, 0, sarama.OffsetOldest)
	if err != nil {
		panic(err)
	}

	Consumer = consumer

	go func() {
		for {
			select {

			case err := <-consumer.Errors():
				log.Fatal(err)

			case m := <-consumer.Messages():
				message := []byte(m.Value)

				jsonMessage := make(map[string]string, 0)
				err := json.Unmarshal(message, &jsonMessage)

				if err != nil {
					//TODO: error
				}

				if sprintID, ok := jsonMessage["sprintID"]; ok {
					sprintUUI, err := gocql.ParseUUID(sprintID)
					if err != nil {
						//TODO: error
					}

					if hub, ok := ActiveHubs[sprintUUI]; ok {
						hub.Broadcast <- message
					}
				}
			}
		}
	}()
}
