
package main

import (
    "fmt"
    "log"
    "math/rand"
    "time"
    "github.com/Shopify/sarama"
)

func main() {
    // Configuring Kafka producer
    config := sarama.NewConfig()
    config.Producer.RequiredAcks = sarama.WaitForAll
    config.Producer.Retry.Max = 5
    config.Producer.Return.Successes = true

    // Connect to Kafka broker
    producer, err := sarama.NewSyncProducer([]string{"localhost:9092"}, config)
    if err != nil {
        log.Fatal("Failed to connect to Kafka broker: ", err)
    }
    defer producer.Close()

    ticker := time.NewTicker(2 * time.Second) // Generate new stock prices every 2 seconds
    for range ticker.C {
        price := rand.Float64() * 100
        stockSymbol := "AAPL" // For now, using a single stock symbol. You can later add more.
        msg := &sarama.ProducerMessage{
            Topic: "stock-prices",
            Value: sarama.StringEncoder(fmt.Sprintf("%s: %.2f", stockSymbol, price)),
        }

        partition, offset, err := producer.SendMessage(msg)
        if err != nil {
            log.Printf("Failed to send message: %s", err)
        }

        log.Printf("Stock price published: %s $%.2f, partition: %d, offset: %d", stockSymbol, price, partition, offset)
    }
}
