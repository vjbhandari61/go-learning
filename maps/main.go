package main

import "fmt"

func main() {
	ages := map[string]int{
		"Alice":   30,
		"Bob":     25,
		"Charlie": 35,
	}

	ages["Charlie"] = 36
	fmt.Println("Bob's age:", ages["Bob"])
}
