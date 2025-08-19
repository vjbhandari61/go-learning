package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Contact struct {
	Name  string
	Phone string
}

func main() {
	contacts := []Contact{}
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Welcome To The PhoneBook")
		fmt.Println("1. Add Contact")
		fmt.Println("2. List Contact")
		fmt.Println("3. Search Contact")
		fmt.Println("4. Exit")
		fmt.Print("Choose an option")

		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			fmt.Println("Enter Name")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)

			fmt.Println("Enter Phone Number")
			phone, _ := reader.ReadString('\n')
			phone = strings.TrimSpace(phone)

			contacts = append(contacts, Contact{Name: name, Phone: phone})
			fmt.Println("Contact Added")

		case "2":
			fmt.Println("All Contacts:")
			for _, contact := range contacts {
				fmt.Printf("Name: %s, Phone: %s\n", contact.Name, contact.Phone)
			}

		case "3":
			fmt.Println("Enter Name to Search")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)

			found := false
			for _, c := range contacts {
				if strings.EqualFold(c.Name, name) {
					fmt.Println("Contact Found:", c.Name, c.Phone)
					found = true
				}
			}
			if !found {
				fmt.Println("Contact Not Found")
			}

		case "4":
			fmt.Println("Exiting PhoneBook")
			return

		default:
			fmt.Println("Invalid Option, Please Try Again")
		}
	}
}
