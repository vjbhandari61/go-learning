package main

import "fmt"

func main() {
	// fixed size array
	var nums [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println("Array = ", nums)

	// dynamic size array
	var fruits []string = []string{"apple", "banana"}
	fruits = append(fruits, "cherry", "date")
	fmt.Println("Fruits = ", fruits)

}
