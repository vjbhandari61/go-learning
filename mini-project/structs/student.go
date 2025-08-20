package structs

import (
	"errors"
	"fmt"
)

type Student struct {
	Name  string
	Age   int
	Grade float64
}

type Person interface {
	Introduce() string
}

func (s Student) Introduce() string {
	return fmt.Sprintf("Hi, I am %s, Grade: %.2f", s.Name, s.Grade)
}

func AverageAge(students []Student) (float64, error) {
	if len(students) == 0 {
		return 0, errors.New("no students found")
	}
	total := 0
	for _, s := range students {
		total += s.Age
	}
	return float64(total) / float64(len(students)), nil
}

func FindStudent(students []Student, name string) (Student, error) {
	for _, s := range students {
		if s.Name == name {
			return s, nil
		}
	}
	return Student{}, errors.New("student not found")
}

func TopPerformer(students []Student) (Student, error) {
	if len(students) == 0 {
		return Student{}, errors.New("no students found")
	}
	top := students[0]
	for _, s := range students {
		if s.Grade > top.Grade {
			top = s
		}
	}
	return top, nil
}
