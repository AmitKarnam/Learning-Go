package main

import "fmt"

type geometry interface {
	area() float64
	perimeter() float64
}

type rect struct {
	widht, height float64
}

/*
(r rect) : This is the receiver of the method, which specifies the type that the method belongs to. In this case, (r rect) indicates that the method area() belongs to the rect type.

func (r rect) area() float64 defines a method named area() that belongs to the rect type and returns a float64 value.
*/
func (r rect) area() float64 {
	return r.widht * r.height
}

func (r rect) perimeter() float64 {
	return 2*r.widht + 2*r.height
}

/*
	=> The measure function takes a parameter g of type geometry, which is an interface. This means that measure can accept any value that satisfies the geometry interface.
	=> The geometry interface specifies two methods: area() and perimeter(). Any type that implements these methods is considered to satisfy the geometry interface.
	=> In the main function, a variable r of type rect is created, which is a struct representing a rectangle with width and height properties.
	=> When measure(r) is called, r is passed as an argument to the measure function. Since rect implements the geometry interface (as it has the area() and perimeter() methods defined), it can be used as an argument for a parameter of type geometry.
	=> Inside the measure function, the g parameter of type geometry is used. Even though the argument passed is of type rect, it is treated as a value of type geometry due to interface satisfaction.

*/

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perimeter())
}

func main() {
	r := rect{widht: 3, height: 4}

	measure(r)
}
