package main

import "fmt"

func reverseGivenInteger(num int) int {
	reservedNum := 0
	for num > 0 {
		remainder := num % 10
		reservedNum *= 10
		reservedNum += remainder
		num /= 10
	}
	return reservedNum
}

func main() {
	var num = 3849
	fmt.Printf("Reversed integer is: %d\n", reverseGivenInteger(num))
	num = 2222
	fmt.Printf("Reversed integer is: %d\n", reverseGivenInteger(num))
	num = 1020
	fmt.Printf("Reversed integer is: %d\n", reverseGivenInteger(num))
	num = 6031
	fmt.Printf("Reversed integer is: %d\n", reverseGivenInteger(num))
}
