package main

import (
	"fmt"
	"log"
	"time"
)

type Block struct {
	nonce        int
	previousHash string
	timestamp    int64
	transactions []string
}

/*
b := new(Block)
b.timestamp = time.Now().UnixNano()
return b
*/
func NewBlock(nonce int, previousHash string) *Block {
	return &Block{
		nonce:        nonce,
		previousHash: previousHash,
		timestamp:    time.Now().UnixNano(),
	}
}

func (b *Block) Print() {
	fmt.Printf("timestamp			%d\n", b.timestamp)
	fmt.Printf("nonce				%d\n", b.nonce)
	fmt.Printf("previous_hash			%s\n", b.previousHash)
	fmt.Printf("transactions			%s\n", b.transactions)
}

func main() {
	log.SetPrefix("BLOCKNET: ")

	b := NewBlock(0, "init hash")
	b.Print()
}
