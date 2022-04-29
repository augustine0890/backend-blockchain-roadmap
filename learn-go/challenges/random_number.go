package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randomNumberGeneratorInRange(rangeStart, rangeEnd int) int {
	rand.Seed(time.Now().UnixNano())
	if rangeStart < rangeEnd {
		return rand.Intn(rangeEnd-rangeStart+1) + rangeStart
	}
	return -1
}

func main() {
	fmt.Printf("My random number: %v\n", randomNumberGeneratorInRange(10, 50))
	fmt.Printf("My random number: %v\n", randomNumberGeneratorInRange(100, 200))
	fmt.Printf("My random number: %v\n", randomNumberGeneratorInRange(40, 20))
}
