package main

import "fmt"

func main() {
	var m, n int = 20, 30
	var a, b string = "abc", "def"
	var x, y, z, t, r int
	var str string

	x = m + n
	fmt.Println(x)
	y = n - m
	fmt.Println(y)
	z = m * n
	fmt.Println(z)
	t = n / m
	fmt.Println(t)
	r = n % m
	fmt.Println(r)
	str = a + " " + b
	fmt.Println(str)
}
