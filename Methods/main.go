package main

import "fmt"

type Rectangle struct {
	width  float64
	height float64
}

// function with value receiver
func (r Rectangle) Area() float64 {
	return r.height * r.width
}

// function with value receiver
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

// function  with pointer receiver
func (r *Rectangle) Scale(factor float64) {
	r.width *= factor
	r.height *= factor
}

func main() {
	rect := Rectangle{width: 10, height: 15}

	area := rect.Area()
	perimter := rect.Perimeter()

	fmt.Println("Area of the rectangle : ", area)
	fmt.Println("Perimeter of the rectangle : ", perimter)

	rect.Scale(2)
	fmt.Println("\nNew Dimensions after scale")
	fmt.Println("Width : ", rect.width, "Height : ", rect.height)
}
