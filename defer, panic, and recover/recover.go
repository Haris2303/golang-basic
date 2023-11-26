package main

import "fmt"

func main() {
	runApp(true)
	fmt.Println("Hello Bang")
}

func endApp() {
	fmt.Println("End App")
	message := recover()
	fmt.Println("Terjadi panic", message)
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("Ups Error")
	}
}
