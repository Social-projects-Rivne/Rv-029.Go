package scrum_poker

import (
	"github.com/Shopify/sarama"
)

func InitConsumer(topic string, partition int32) sarama.PartitionConsumer {

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

	return consumer

}
