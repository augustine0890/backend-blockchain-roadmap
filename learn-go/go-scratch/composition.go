package main

import (
	"fmt"
	"log"
	"os"
)

type User struct {
	Id             int
	Name, Location string
}

func (u *User) Greetings() string {
	return fmt.Sprintf("Hi %s from %s", u.Name, u.Location)
}

type Player struct {
	u      User
	GameId int
}

func GreetingsForPlayer(p Player) string {
	return p.u.Greetings()
}

type Job struct {
	Commad string
	Logger *log.Logger
}

func main() {
	p := Player{}
	p.u.Id = 42
	p.u.Name = "Matt"
	p.u.Location = "LA"
	fmt.Println(p.u.Greetings())
	fmt.Println("ForPlayer:", GreetingsForPlayer(p))

	job := &Job{"demo", log.New(os.Stdout, "Job: ", log.Ldate)}
	job.Logger.Print("starting now...")
}
