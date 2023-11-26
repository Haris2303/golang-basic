package main

import "fmt"

func main() {
	var a = 10
	b := 10
	c := a + b

	fmt.Println(c)

	a += 10
	fmt.Println(a)

	b++
	fmt.Println(b)
}
