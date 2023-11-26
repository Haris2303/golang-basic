package main

import "fmt"

func main() {
	sayHelloWithFilter("Ucup", spamFilter)

	filter := spamFilter
	sayHelloWithFilter("Taek", filter)
}

type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) {
	filteredName := filter(name)
	fmt.Println("Hello", filteredName)
}

func spamFilter(name string) string {
	if name == "Taek" {
		return "..."
	} else {
		return name
	}
}
