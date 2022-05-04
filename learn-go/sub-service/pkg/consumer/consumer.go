package consumer

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/Shopify/sarama"
)

type ConsumerGroupHandler interface {
	sarama.ConsumerGroupHandler
	WaitReady()
	Reset()
}

type ConsumerGroup struct {
	cg sarama.ConsumerGroup
}

func NewConsumerGroup(broker string, topics []string, group string, handler ConsumerGroupHandler) (*ConsumerGroup, error) {
	ctx := context.Background()
	cfg := sarama.NewConfig()
	cfg.Consumer.Offsets.Initial = sarama.OffsetOldest
	client, err := sarama.NewConsumerGroup([]string{broker}, group, cfg)
	if err != nil {
		panic(err)
	}

	go func() {
		for {
			err := client.Consume(ctx, topics, handler)
			if err != nil {
				if err == sarama.ErrClosedConsumerGroup {
					break
				} else {
					panic(err)
				}
			}
			if ctx.Err() != nil {
				return
			}
			handler.Reset()
		}
	}()

	// Await till the consumer has been set up
	handler.WaitReady()

	return &ConsumerGroup{
		cg: client,
	}, nil
}

func (c *ConsumerGroup) Close() error {
	return c.cg.Close()
}

type ConsumerSessionMessage struct {
	Session  sarama.ConsumerGroupSession
	Messages *sarama.ConsumerMessage
}

func decodeMessage(data []byte) error {
	var msg Messages
	err := json.Unmarshal(data, &msg)
	if err != nil {
		return err
	}
	return nil
}

func StartSyncConsumer(broker, topic string) (*ConsumerGroup, error) {
	var count int64
	var start = time.Now()
	handler := NewSyncConsumerGroupHandler(func(data []byte) error {
		if err := decodeMessage(data); err != nil {
			return err
		}
		fmt.Printf("Messages: %v\n", string(data))
		count++
		if count%100 == 0 {
			fmt.Printf("sync consumer consumed %d message at speed %.2f/s\n", count, float64(count)/time.Since(start).Seconds())
		}
		return nil
	})

	consumer, err := NewConsumerGroup(broker, []string{topic}, "sync-consumer-"+fmt.Sprintf("%d", time.Now().Unix()), handler)
	if err != nil {
		return nil, err
	}
	return consumer, nil
}
