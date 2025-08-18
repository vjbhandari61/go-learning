package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func divide(a int, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero cannot be performed")
	}
	return a / b, nil
}

func main() {
	sum := add(3, 5)
	fmt.Println("The sum is:", sum)

	quotient, err := divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("The quotient is:", quotient)
	}
}
