package main

import (
	"fmt"
	"strings"
)

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}

var names = []string{"Katrina", "Evan", "Neil", "Adam", "Martin", "Matt",
	"Emma", "Isabella", "Emily", "Madison",
	"Ava", "Olivia", "Sophia", "Abigail",
	"Elizabeth", "Chloe", "Samantha",
	"Addison", "Natalie", "Mia", "Alexis",
}

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	count := map[string]int{}
	for _, w := range words {
		count[w]++
	}
	return count
}

func main() {
	m["Splice"] = Vertex{34.05641, -118.48175}
	fmt.Println(m["Splice"])
	delete(m, "Splice")
	fmt.Printf("%v\n", m)
	name, ok := m["Splice"]
	fmt.Printf("key 'Splice is present?: %t - value: %v\n", ok, name)
	name, ok = m["Google"]
	fmt.Printf("key 'Google is present?: %t - value: %v\n", ok, name)

	var maxLen int
	for _, name := range names {
		if l := len(name); l > maxLen {
			maxLen = l
		}
	}
	output := make([][]string, maxLen)

	for _, name := range names {
		output[len(name)-1] = append(output[len(name)-1], name)
	}
	fmt.Println(output)

	count := WordCount("I ate a donut")
	fmt.Println(count)
}
