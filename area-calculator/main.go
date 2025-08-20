package main

import (
	"fmt"
	"shapes/shapes"
)

func calculateArea(s shapes.Shape) float64 {
	return s.Area()
}

func main() {
	c := shapes.Circle{Radius: 5}
	r := shapes.Rectangle{Length: 10, Breadth: 5}

	fmt.Println("Area of a Rectangle", calculateArea(r))
	fmt.Println("Area of a Circle", calculateArea(c))
}
