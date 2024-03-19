package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	// // pass by value
	// address1 := Address{"Sorong", "Manokwari", "Jayapura"}
	// address2 := address1

	// address2.City = "Ambon"
	// fmt.Println(address1) // tidak berubah
	// fmt.Println(address2) // berubah menjadi bandung

	// ============

	// pass by reference (Pointer)
	var address1 Address = Address{"Sorong", "Papua Barat Daya", "Indonesia"}
	var address2 *Address = &address1

	address2.City = "Ambon"
	fmt.Println(address1) // ikut berubah
	fmt.Println(address2) // berubah menjadi bandung
}
