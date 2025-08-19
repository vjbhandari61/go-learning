package main

import "fmt"

type Contact struct {
	Name  string
	Phone string
}

func main() {
	contact := Contact{
		Name:  "Vijay",
		Phone: "123-456-7890",
	}

	fmt.Println("Contact Name:", contact.Name, "\nPhone:", contact.Phone)
}
