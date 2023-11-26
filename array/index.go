package main

import "fmt"

func main() {
	var names [3]string

	names[0] = "Ucup"
	names[1] = "Asep"
	names[2] = "Otong"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [3]int{
		90,
		80,
		70,
	}

	fmt.Println(values)
	fmt.Println(len(values))

	values[2] = 75
	fmt.Println(values[2])
}
