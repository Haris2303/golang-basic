package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	var address1 Address = Address{"Sorong", "Papua Barat Daya", "Indonesia"}
	var address2 *Address = &address1
	address2.City = "Ambon"
	fmt.Println(address1) // ikut berubah
	fmt.Println(address2) // berubah menjadi bandung

	// address2 = &Address{"Jakarta", "DKI Jakarta", "Indonesia"}
	// fmt.Println(address1) // tidak berubah
	// fmt.Println(address2)

	*address2 = Address{"Jakarta", "DKI Jakarta", "Indonesia"}
	fmt.Println(address1) // ikut berubah
	fmt.Println(address2)
}
