package main

import "fmt"

func main() {
	x := 5

	// Simple if statement
	if x > 0 {
		fmt.Println(("x is positive"))
	}

	// If-else statement
	if x < 0 {
		fmt.Println("x is negative")
	} else {
		fmt.Println("x is positive")
	}

	// If-else if-else statement
	if x < 0 {
		fmt.Println("x is negative")
	} else if x == 0 {
		fmt.Println("x is zero")
	} else {
		fmt.Println("x is positive")
	}

	// Nested if statement
	if x > 0 {
		fmt.Println("x is positive")
		if x > 10 {
			fmt.Println("x is greater than 10")
		} else {
			fmt.Println("x is not greater than 10")
		}
	} else {
		fmt.Println("x is not positive")
	}

	// Switch statement
	switch x {
	case 5:
		fmt.Println("x is five")
	default:
		fmt.Println("x is not five")
	}

}
