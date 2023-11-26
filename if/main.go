package main

import "fmt"

func main() {
	name := "Otong"

	if name == "Ilham" {
		fmt.Println("Hello Ilham")
	} else if name == "Otong" {
		fmt.Println("Hello Otong")
	} else {
		fmt.Println("Kamu Siapa?")
	}

	if length := len(name); length > 5 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama Anda keren")
	}
}
