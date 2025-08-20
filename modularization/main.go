package main

import (
	"fmt"
	"modularization/mathutil"
)

func main() {
	sum := mathutil.Add(5, 3)
	diff := mathutil.Subtract(10, 4)

	fmt.Println("Sum:", sum)
	fmt.Println("Difference:", diff)
}
