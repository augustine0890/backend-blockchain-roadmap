package main

import "fmt"

func main() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)

	cites := make([]string, 3)
	cites[0] = "Santa Monica"
	cites[1] = "Venice"
	cites[2] = "Los Angeles"
	fmt.Printf("%q\n", cites)

	newCities := []string{}
	fmt.Println(len(newCities))
	newCities = append(newCities, "San Diego", "Mountain View")
	fmt.Println(newCities)

	cites = append(cites, newCities...)
	fmt.Printf("%q\n", cites)
	fmt.Println(len(cites))

	countries := make([]string, 42)
	fmt.Println(len(countries))
}
