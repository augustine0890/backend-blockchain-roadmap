package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	// Time out
	response := make(chan *http.Response, 1)
	errors := make(chan *error)

	go func() {
		resp, err := http.Get("https://golangexample.com/")
		if err != nil {
			errors <- &err
		}
		response <- resp
	}()

	for {
		select {
		case r := <-response:
			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("%s\n", string(body[:200]))
			return
		case err := <-errors:
			log.Fatal(*err)
		case <-time.After(1 * time.Second):
			fmt.Println("Timed out!")
			return
		}
	}
}
