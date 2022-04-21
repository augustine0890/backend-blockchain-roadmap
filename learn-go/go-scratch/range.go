package main

import "fmt"

func pow() {
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}

func powShift() {
	pow := make([]int, 10)
	for i := range pow {
		if i%2 == 0 {
			continue // skip each even index of pow
		}
		pow[i] = 1 << uint(i)
		if pow[i] >= 128 {
			break // stop iterating
		}
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}

func main() {
	pow()
	powShift()

	cities := map[string]int{
		"New York":    8336697,
		"Los Angeles": 3857799,
		"Chicago":     2714857,
	}
	for key, value := range cities {
		fmt.Printf("%s has %d inhabitants\n", key, value)
	}
}
