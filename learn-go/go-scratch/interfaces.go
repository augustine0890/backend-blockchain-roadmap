package main

import (
	"fmt"
	"math"
	"os"
	"time"
)

type User struct {
	FirstName, LastName string
}

func (u *User) Name() string {
	return fmt.Sprintf("%s %s", u.FirstName, u.LastName)
}

type Customer struct {
	Id       int
	FullName string
}

func (c *Customer) Name() string {
	return c.FullName
}

type Namer interface {
	Name() string //Both Name() methods can be called using the Namer interface
}

func Greet(n Namer) string {
	return fmt.Sprintf("Dear %s", n.Name())
}

type Reader interface {
	Read(b []byte) (n int, err error)
}

type Writer interface {
	Write(b []byte) (n int, err error)
}

type ReadWriter interface {
	Reader
	Writer
}

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %g", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	return math.Sqrt(x), nil
}

func main() {
	u := &User{"Matt", "Animometti"}
	fmt.Println(Greet(u))
	c := &Customer{42, "Francesc"}
	fmt.Println(Greet(c))

	var w Writer
	// os.Stdout implements Writer
	w = os.Stdout
	fmt.Fprintf(w, "hello, writer\n")

	if err := run(); err != nil {
		fmt.Println(err)
	}
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
