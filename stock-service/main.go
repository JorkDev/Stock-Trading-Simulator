package main

import (
    "github.com/Shopify/sarama"
    "log"
    "math/rand"
    "time"
)

func main() {
    config := sarama.NewConfig()
    config.Producer.RequiredAcks = sarama.WaitForAll
    config.Producer.Retry.Max = 5
    config.Producer.Return.Successes = true

    producer, err := sarama.NewSyncProducer([]string{"localhost:9092"}, config)
    if err != nil {
        panic(err)
    }
    defer producer.Close()

    ticker := time.NewTicker(2 * time.Second)
    for range ticker.C {
        price := rand.Float64() * 100
        msg := &sarama.ProducerMessage{
            Topic: "stock-prices",
            Value: sarama.StringEncoder(fmt.Sprintf("Stock Price: $%.2f", price)),
        }
        producer.SendMessage(msg)
        log.Println("Sent stock price update: $", price)
    }
}
