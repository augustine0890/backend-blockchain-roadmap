package main

import "fmt"

func main() {
	a := [...]string{"hello", "world!"}
	fmt.Println(a)
	fmt.Printf("%s\n", a)
	// Print each element quoted
	fmt.Printf("%q\n", a)

	var x [2][3]string
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			x[i][j] = fmt.Sprintf("row %d - column %d", i+1, j+1)
		}
	}
	fmt.Printf("%q\n", x)
}
