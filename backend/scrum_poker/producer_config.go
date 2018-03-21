package scrum_poker

import "github.com/Shopify/sarama"

func InitProducer() sarama.SyncProducer {
	config := sarama.NewConfig()
	config.Producer.Retry.Max = 10
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true

	producer, err := sarama.NewSyncProducer([]string{"127.0.0.1:9092"}, config)
	if err != nil {
		panic(err)
	}

	return producer
}
