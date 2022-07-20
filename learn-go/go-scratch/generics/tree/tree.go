package main

import (
	"fmt"
	"strings"
)

type Orderable[T any] interface {
	Order(T) int
}

type Tree[T Orderable[T]] struct {
	val         T
	left, right *Tree[T]
}

func (t *Tree[T]) Insert(val T) *Tree[T] {
	if t == nil {
		return &Tree[T]{val: val}
	}

	switch comp := val.Order(t.val); {
	case comp < 0:
		t.left = t.left.Insert(val)
	case comp > 0:
		t.right = t.right.Insert(val)
	}
	return t
}

func (t *Tree[T]) Contains(val T) bool {
	if t == nil {
		return false
	}
	switch comp := val.Order(t.val); {
	case comp < 0:
		return t.left.Contains(val)
	case comp > 0:
		return t.right.Contains(val)
	default:
		return true
	}
}

type OrderableInt int

func (oi OrderableInt) Order(val OrderableInt) int {
	return int(oi - val)
}

type OrderableString string

func (os OrderableString) Order(val OrderableString) int {
	return strings.Compare(string(os), string(val))
}

func main() {
	var it *Tree[OrderableInt]
	it = it.Insert(OrderableInt(5))
	it = it.Insert(OrderableInt(3))
	fmt.Println(it.Contains(2))
	fmt.Println(it.Contains(3))
	a := 10
	it = it.Insert(OrderableInt(a))
	// uncomment to see a compile-time error
	// it = it.Insert(OrderableString("not compile"))
}
