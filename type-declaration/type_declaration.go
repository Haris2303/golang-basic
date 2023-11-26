package main

import "fmt"

func main() {
	type Nim string

	var myNim Nim = "20022241"

	var contoh string = "321343432"
	var contohNim Nim = Nim(contoh)

	fmt.Println(myNim)
	fmt.Println(contohNim)
}
