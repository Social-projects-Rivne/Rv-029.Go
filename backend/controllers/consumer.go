package controllers

import (
	"github.com/Shopify/sarama"
)

var Consumer sarama.PartitionConsumer
var Master sarama.Consumer

func KafkaConsumerInit() {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	brokers := []string{"127.0.0.1:9092"}
	topic := "test-topic-1"

	master, err := sarama.NewConsumer(brokers, config)
	if err != nil {
		panic(err)
	}

	Master = master

	consumer, err := master.ConsumePartition(topic, 0, sarama.OffsetOldest)
	if err != nil {
		panic(err)
	}

	Consumer = consumer
}
