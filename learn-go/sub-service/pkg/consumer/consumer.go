package consumer

import (
	"strings"
	"sub-service/pkg/nft"
	"sync"

	"github.com/Shopify/sarama"
	log "github.com/sirupsen/logrus"
)

type Consumer struct {
	Polygon *nft.NFT
}

func NewConsumer(polygon *nft.NFT) *Consumer {
	return &Consumer{
		Polygon: polygon,
	}
}

func (c *Consumer) StartConsumer() {
	broker := "13.92.179.244:9092"
	topics := []string{"MATIC-PDMP-SERVICE"}

	log.Info("[🚀 Consumer is running]")

	cfg := sarama.NewConfig()
	cfg.Consumer.Offsets.Initial = sarama.OffsetOldest

	// Create new sarama consumer
	master, err := sarama.NewConsumer([]string{broker}, cfg)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := master.Close(); err != nil {
			panic(err)
		}
	}()

	errors := make(chan *sarama.ConsumerError)
	for _, topic := range topics {
		if strings.Contains(topic, "__consumer_offsets") {
			continue
		}
		partitions, _ := master.Partitions(topic)
		// this only consumes partition no 1, you would probably want to consume all partitions
		consumer, err := master.ConsumePartition(topic, partitions[0], sarama.OffsetOldest)
		if nil != err {
			// log.Info("Topic %v Partitions: %v", topic, partitions)
			panic(err)
		}

		var wg sync.WaitGroup
		wg.Add(1)
		go func(topic string, consumer sarama.PartitionConsumer) {
			defer wg.Done()
			for {
				select {
				case consumerError := <-consumer.Errors():
					errors <- consumerError
					log.Info("consumerError: ", consumerError.Err)
				case msg := <-consumer.Messages():
					if msg.Topic == topic {
						log.Infof("TOPIC: %s - OFFSET: %d - KEY: %s - MESSAGE: %s", msg.Topic, msg.Offset, string(msg.Key), string(msg.Value))
					}
					// fmt.Println("Got message on topic ", topic, msg.Value)
				}
			}
		}(topic, consumer)
		wg.Wait()
	}
}
