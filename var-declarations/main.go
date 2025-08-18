package main

import "fmt"

func main() {
	// Statically typed variable declarations
	var name string = "Vijay Bhandari"
	var age int = 20
	var isEmployed bool = true

	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Is Employed:", isEmployed)

	// Dynamic variable declarations
	name2 := "Arun Kumar"
	age2 := 25
	isEmployed2 := false

	fmt.Println("\nName2:", name2)
	fmt.Println("Age2:", age2)
	fmt.Println("Is Employed2:", isEmployed2)

	// Constants declarations
	const pi float64 = 3.14

	fmt.Println("\nValue of Pi:", pi)
}
