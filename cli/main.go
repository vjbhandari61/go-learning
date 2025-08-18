package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter a name: ")

	name, _ := reader.ReadString('\n')

	fmt.Println("\nHello,", name)
	fmt.Println("Welcome to the Go programming language!")
}
