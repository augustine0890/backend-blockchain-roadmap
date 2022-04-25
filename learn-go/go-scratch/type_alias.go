package main

import (
	"fmt"
	"math"
	"strings"
)

type MyStr string

func (s MyStr) Uppercase() string {
	return strings.ToUpper(string(s))
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	fmt.Println(MyStr("test").Uppercase())

	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}
