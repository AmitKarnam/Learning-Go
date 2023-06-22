package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {

	// Declaring a variable of type person
	var p1 person = person{"Sam", 21}

	// The fields of a struct are accessed by a dot(.) operator
	fmt.Println(p1.name)
	fmt.Println(p1.age)

	fmt.Println("An Array of type person ")

	// Declaring an array of struct people
	var p_arr []person = []person{
		{"John", 30},
		{"Jake", 25},
		{"Pearson", 23},
	}

	// Iterating through the array of person struct
	for _, people := range p_arr {
		fmt.Println(people.name)
		fmt.Println(people.age)
	}

	// Pointer to a struct

	var pointer *person = &p1

	// The star(*) operator is used to access the value stored in a pointer, just like C programming language syntax.

	fmt.Println((*pointer))
}
