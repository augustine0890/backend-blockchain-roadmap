package main

import "fmt"

func GetSum(array []int) int {
	sum := 0
	for i := 0; i < len(array); i++ {
		sum += array[i]
	}
	return sum
}

var (
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie",
		"Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distribution = make(map[string]int, len(users))
)

func main() {
	fmt.Println("get sum: ", GetSum([]int{1, 2, 3, 4}))

	for _, user := range users {
		var total int

		for _, c := range user {
			switch string(c) {
			case "a", "A":
				total++
			case "e", "E":
				total++
			case "i", "I":
				total += 2
			case "o", "O":
				total += 3
			case "u", "U":
				total += 4
			}
		}
		if total > 10 {
			total = 10
		}
		distribution[user] = total
		coins = coins - total
	}

	fmt.Println(distribution)
	fmt.Println("Coins left: ", coins)
}
