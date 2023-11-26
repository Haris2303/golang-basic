package main

import "fmt"

func main() {
	// fisrtname, lastname := getFullName()
	// fmt.Println(fisrtname, lastname)

	fisrtname, _ := getFullName()
	fmt.Println(fisrtname)
}

func getFullName() (string, string) {
	return "Ucup", "Surucup"
}
