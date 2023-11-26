package main

import "fmt"

func main() {
	const firstName string = "Ucup"
	const lastName = "Surucup"

	fmt.Println(firstName)
	// fmt.Println(lastName)

	// assignment contant sekaligus
	const (
		product = "Handphone"
		price   = 5000000
	)

	fmt.Println("Product = ", product)
}
