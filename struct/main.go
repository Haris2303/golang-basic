package main

import "fmt"

// Cara 1
type Customer struct {
	Name, Address string
	Age           int
}

func main() {
	var customer1 Customer
	fmt.Println(customer1)

	customer1.Name = "Ucup Surucup"
	customer1.Address = "Jalan Ucup"
	customer1.Age = 20
	fmt.Println(customer1)
	fmt.Println(customer1.Name)
	fmt.Println(customer1.Address)
	fmt.Println(customer1.Age)

	// Cara 2
	customer2 := Customer{
		Name:    "Otong",
		Address: "Singapura",
		Age:     21,
	}
	fmt.Println(customer2)

	// cara 3
	customer3 := Customer{"Udin", "Thailand", 32}
	fmt.Println(customer3)
}
