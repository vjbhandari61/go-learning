package main

import "fmt"

func update(x *int) {
	*x = *x + 1
}

func main() {
	n := 10
	update(&n)
	fmt.Println(n)
}
