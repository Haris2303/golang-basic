package main

import "fmt"

func main() {
	var i int

	for i = 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println("Hello ke", i)
	}
}
