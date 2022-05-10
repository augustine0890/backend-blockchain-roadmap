package main

import (
	"fmt"
	"os"
	"os/signal"
	"sub-service/pkg/consumer"
	"sub-service/pkg/nft"
)

func main() {
	// var p *nft.NFT
	p := nft.NewNFT()
	kafka := consumer.NewConsumer(p)
	go kafka.StartConsumer()
	p.ProcessMsg()
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	// Get signnal for finish
	doneCh := make(chan struct{})
	go func() {
		for {
			select {
			case <-signals:
				fmt.Println("Interrupt is detected")
				doneCh <- struct{}{}
			}
		}
	}()

	<-doneCh
}
