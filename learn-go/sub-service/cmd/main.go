package main

import (
	"fmt"
	"os"
	"os/signal"
	"sub-service/pkg/consumer"
	"syscall"
)

const SyncMode = "sync"

func main() {
	broker := "13.92.179.244:9092"
	topic := "MATIC-PDMP-SERVICE"
	var done = make(chan struct{})
	defer close(done)
	mode := "sync"

	switch mode {
	case SyncMode:
		consumer, err := consumer.StartSyncConsumer(broker, topic)
		if err != nil {
			panic(err)
		}
		defer consumer.Close()
	default:
		return
	}

	// Set up a channel for handling Ctrl-C, etc
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	fmt.Println("received signal", <-sigchan)
}
