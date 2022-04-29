package main

import (
	"fmt"
)

var s = "Golang is awesome"

func reverseAString(s string) string {
	var reverved string
	for i := 0; i < len(s); i++ {
		reverved = string(s[i]) + reverved
	}
	return reverved
}

func reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}

func main() {
	fmt.Printf("Reversed string is: %s\n", reverseAString(s))
	fmt.Printf("Reversed string is: %s\n", reverse(s))

}
