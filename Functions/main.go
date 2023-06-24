package main

import "fmt"

// custom function
func add(a int, b int) int {
	return (a + b)
}

func init() {
	fmt.Println("This is an init function.")
	// Do some initialisation
}

// closure function

func outerFunc(s string) func() {
	message := "Hello"

	closure := func() {
		fmt.Println(message + " " + s)
	}

	return closure
}

// Main function
func main() {
	var res int = add(1, 2)
	fmt.Println("Addition of 2 nos : ", res)

	// Executing the closure function
	fmt.Println("The result of the closure funciton")
	myclosure := outerFunc("James")

	myclosure()
}
