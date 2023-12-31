package main

import "fmt"

func main() {
	fmt.Println(factorialLoop(5))
	fmt.Println(factorialRecursive(5))
}

func factorialLoop(value int) int {
	result := 1

	for i := value; i > 0; i-- {
		result *= i
	}

	return result
}

func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialLoop(value-1)
	}
}
