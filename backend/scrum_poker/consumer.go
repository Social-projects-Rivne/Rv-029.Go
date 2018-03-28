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

				jsonMessage := make(map[string]interface{}, 0)
				err := json.Unmarshal(message, &jsonMessage)
				if err != nil {
					//TODO: error
				}

				if action, ok := jsonMessage["action"]; ok && action == "ESTIMATION" {
					if issueID, ok := jsonMessage["issue_uuid"]; ok {
						issueUUID, err := gocql.ParseUUID(issueID.(string))
						if err != nil {
							log.Println("Invalid issue id received from kafka")
						}

						if hub, ok := ActiveHubs[issueUUID]; ok {
							userUUID, err := gocql.ParseUUID(jsonMessage["user_uuid"].(string))
							if err != nil {
								log.Println("Invalid user id received from kafka")
							}
							hub.Summary[userUUID] = int(jsonMessage["estimate"].(float64))

							hub.Calculate()
						}
					}
				}
			}
		}
	}()
}
