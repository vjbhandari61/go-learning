package main

import "fmt"

type Rectangle struct {
	Length  int
	Breadth int
}

func (r Rectangle) Area() int {
	return r.Length * r.Breadth
}

func main() {
	rect := Rectangle{Length: 10, Breadth: 20}
	fmt.Println("Area of this rectangle =", rect.Area())
}
