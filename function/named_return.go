package main

import "fmt"

func main() {
	a, b, c := getCompleteName()
	fmt.Println(a, b, c)
}

func getCompleteName() (firstname, middlename, lastname string) {
	firstname = "Ucup"
	middlename = "Surucup"
	lastname = "Markucup"

	return firstname, middlename, lastname
}
