package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"sub-service/pkg/consumer"

	"github.com/Shopify/sarama"
)

func main() {
	broker := "13.92.179.244:9092"
	topic := "MATIC-PDMP-SERVICE"

	log.Println("Starting a new Sarama consumer")

	cfg := sarama.NewConfig()
	cfg.Consumer.Offsets.Initial = sarama.OffsetOldest

	// Create new consumer
	master, err := sarama.NewConsumer([]string{broker}, cfg)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := master.Close(); err != nil {
			panic(err)
		}
	}()

	// topics, _ := master.Topics()
	consumer, errors := consumer.StartConsumer([]string{topic}, master)

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	// Count how many message processed
	msgCount := 0

	// Get signnal for finish
	doneCh := make(chan struct{})
	go func() {
		for {
			select {
			case msg := <-consumer:
				msgCount++
				fmt.Println("Received messages", msg.Topic, msg.Offset, string(msg.Key), string(msg.Value))
			case consumerError := <-errors:
				msgCount++
				fmt.Println("Received consumerError ", string(consumerError.Topic), string(consumerError.Partition), consumerError.Err)
				doneCh <- struct{}{}
			case <-signals:
				fmt.Println("Interrupt is detected")
				doneCh <- struct{}{}
			}
		}
	}()

	<-doneCh
	fmt.Println("Processed", msgCount, "messages")
}
