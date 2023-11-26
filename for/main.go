package main

import "fmt"

func main() {
	counter := 1

	for counter <= 10 {
		fmt.Println("Counter ke", counter)
		counter++
	}

	for i := 0; i < 10; i++ {
		fmt.Println("Index ke", i)
	}

	names := []string{"Ucup", "Otong", "Udin"}
	for index, name := range names {
		fmt.Printf("Index ke %d = %s\n", index, name)
	}
}
