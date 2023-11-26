package main

import "fmt"

func main() {
	var name string

	name = "Ucup Surucup"
	fmt.Println(name)

	name = "Otong Surotong"
	fmt.Println(name)

	var umur = 5
	fmt.Println(umur)

	product := "Handphone" // alias : var product = "Handphone"
	fmt.Println("Product =", product)

	product = "Iphone"
	fmt.Println("Product =", product)

	// Assignment sekaligus
	var (
		firstName  = "Ucup"
		middleName = "Surucup"
		lastName   = "Markucup"
	)

	fmt.Println(firstName)
	fmt.Println(middleName)
	fmt.Println(lastName)
}
