package main

import "fmt"

func main() {
	var x int
	x = 10
	fmt.Println(x)
	var y int = 20
	fmt.Println(y)
	z := 30
	fmt.Println(z)
	var Sp, dt, _tm, calc = 10, 33, 49, 53
	fmt.Println(Sp, dt, _tm, calc)

	var (
		a = 5
		b = 10
		c = 15
	)
	fmt.Println(a, b, c)

	var i int = 4
	var s string = "abc"
	fmt.Println(i)
	fmt.Println(s)
}
