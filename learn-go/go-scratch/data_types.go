package main

import (
	"fmt"
	"time"
)

type Stringer interface {
	String() string
}

type fakeString struct {
	content string
}

// function used to implement the Stringer interface
func (s *fakeString) String() string {
	return s.content
}

func printString(value interface{}) {
	switch str := value.(type) {
	case string:
		fmt.Println(str)
	case Stringer:
		fmt.Println(str.String())
	}
}

func timeMap(y interface{}) {
	z, ok := y.(map[string]interface{})
	if ok {
		z["updated_at"] = time.Now()
	}
}

func main() {
	var n1 uint8
	n1 = 200
	fmt.Println(n1)

	var n2 uint16
	n2 = 64300
	fmt.Println(n2)

	var n3 uint32
	n3 = 98733423
	fmt.Println(n3)

	var n4 uint64
	n4 = 279283782735793
	fmt.Println(n4)

	foo := map[string]interface{}{
		"Matt": 42,
	}
	timeMap(foo)
	fmt.Println(foo)

	printString("Hello, Gophers")
	s := &fakeString{"Ceci n'est pas um string"}
	printString(s)
}
