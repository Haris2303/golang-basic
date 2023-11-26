package main

import "fmt"

func main() {
	name := "Otong"

	switch name {
	case "Ucup":
		fmt.Println("Hello Ucup")
	case "Otong":
		fmt.Println("Hello Otong")
	default:
		fmt.Println("Hello Person")
	}

	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama sangat panjang")
	case false:
		fmt.Println("Nama tidak panjang")
	}

	length := len(name)

	switch {
	case length > 10:
		fmt.Println("Nama terlalu panjang")
	case length > 5:
		fmt.Println("Nama cukup panjang")
	default:
		fmt.Println("Nama Anda Oke")
	}
}
