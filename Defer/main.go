package main

import "fmt"

func main() {
	fmt.Println("Start")
	defer fmt.Println("Deferred funciton execution 1")
	defer fmt.Println("Deferred function execution 2")
	fmt.Println("End")
}
