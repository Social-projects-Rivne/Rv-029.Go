package controllers

import (
	"github.com/Shopify/sarama"
)

var Producer sarama.SyncProducer

func KafkaProducerInit() {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true

	producer, err := sarama.NewSyncProducer([]string{"127.0.0.1:9092"}, config)
	if err != nil {
		panic(err)
	}

	Producer = producer
}
