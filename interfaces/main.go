package main

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Length  float64
	Breadth float64
}

func (r Rectangle) Area() float64 {
	return r.Length * r.Breadth
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func printArea(s Shape) {
	println("Area:", s.Area())
}

func main() {
	rect := Rectangle{Length: 10, Breadth: 20}
	circ := Circle{Radius: 5}

	printArea(rect)
	printArea(circ)
}
