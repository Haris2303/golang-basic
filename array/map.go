package main

import "fmt"

func main() {
	// var person map[string]string = map[string]string{}
	// person["name"] = "Otong"
	// person["address"] = "Jalan Santai"

	person := map[string]string{
		"name":    "Ucup",
		"address": "Jalan Bima",
	}

	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person)

	book := make(map[string]string)
	book["title"] = "Belajar Golang"
	book["author"] = "Ucup"
	book["ups"] = "Salah"

	fmt.Println(book)
	fmt.Println(len(book))

	delete(book, "ups")

	fmt.Println(book)

}
