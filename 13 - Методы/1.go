package main

import (
	"fmt"
	"time"
)

type Student struct {
	Name      string
	BirthYear int
	Course    int
	AvgGrade  float64
}

func (s Student) Age() int {
	return time.Now().Year() - s.BirthYear
}

func (s Student) Status() string {
	if s.AvgGrade >= 4.5 {
		return "Отличник"
	} else if s.AvgGrade >= 3.5 {
		return "Хорошист"
	}
	return "Троечник"
}

func main() {
	student := Student{
		Name:      "Вейнмайер Андрей",
		BirthYear: 2005,
		Course:    3,
		AvgGrade:  4.8,
	}

	fmt.Printf("%s: возраст %d, статус %s\n",
		student.Name, student.Age(), student.Status())
}
