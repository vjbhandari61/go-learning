package main

import (
	"errors"
	"fmt"
)

func divide(a int, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by 0")
	}

	return a / b, nil
}

func main() {
	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("Value of a = ", result)
	}
}
