package main

import "fmt"

func main() {
	var int_var int
	int_var = 10

	fmt.Println("Integer Variable : ", int_var)

	var float_num_32bit float32
	var float_num_64bit float64

	float_num_32bit = 23.3422
	float_num_64bit = 9.23940134790

	fmt.Println("Floating Point Variable with 32 bit precision : ", float_num_32bit)
	fmt.Println("Floating Point Variable with 64 bit precision : ", float_num_64bit)

	var str_var string
	str_var = "This is a string."

	fmt.Println("String Variable : ", str_var)

	// Declarationa and Initialization in one statement example

	str_var1 := "This is another string."
	fmt.Println("\nstr_var1 is declared and initialised in one single sentnace in the above example.")
	fmt.Println("str_var1 : ", str_var1)
}
