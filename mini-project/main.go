package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"mini-project/structs"
)

func main() {
	var students []structs.Student
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n1. Add Student")
		fmt.Println("2. List Students")
		fmt.Println("3. Search Student")
		fmt.Println("4. Average Age")
		fmt.Println("5. Top Performer")
		fmt.Println("6. Exit")
		fmt.Print("Choose an option: ")

		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			fmt.Print("Enter name: ")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)

			fmt.Print("Enter age: ")
			ageInput, _ := reader.ReadString('\n')
			ageInput = strings.TrimSpace(ageInput)
			age, _ := strconv.Atoi(ageInput)

			fmt.Print("Enter grade: ")
			gradeInput, _ := reader.ReadString('\n')
			gradeInput = strings.TrimSpace(gradeInput)
			grade, _ := strconv.ParseFloat(gradeInput, 64)

			students = append(students, structs.Student{Name: name, Age: age, Grade: grade})
			fmt.Println("Student added!")

		case "2":
			fmt.Println("\n--- Student List ---")
			for _, s := range students {
				fmt.Println(s.Introduce())
			}

		case "3":
			fmt.Print("Enter name to search: ")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)

			student, err := structs.FindStudent(students, name)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Found:", student.Introduce())
			}

		case "4":
			avg, err := structs.AverageAge(students)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Printf("Average Age: %.2f\n", avg)
			}

		case "5":
			top, err := structs.TopPerformer(students)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Printf("Top Performer: %s with Grade %.2f\n", top.Name, top.Grade)
			}

		case "6":
			fmt.Println("Goodbye!")
			return

		default:
			fmt.Println("Invalid option")
		}
	}
}
