package main

import (
	"fmt"
	"log"

	"github.com/Shopify/sarama"
)

func main() {
	consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, nil)
	if err != nil {
		log.Fatal("Error creating Kafka consumer: ", err)
	}
	defer consumer.Close()

	partitionConsumer, err := consumer.ConsumePartition("stock-prices", 0, sarama.OffsetNewest)
	if err != nil {
		log.Fatal("Error starting partition consumer: ", err)
	}
	defer partitionConsumer.Close()

	for message := range partitionConsumer.Messages() {
		fmt.Printf("Received stock price update: %s\n", string(message.Value))
	}
}
