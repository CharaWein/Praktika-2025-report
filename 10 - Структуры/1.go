package main

import "fmt"

type Student struct {
	Name     string
	Age      int
	Course   int
	AvgGrade float64
}

func (s Student) GetStatus() string {
	if s.AvgGrade >= 4.5 {
		return "Отличник"
	} else if s.AvgGrade >= 3.5 {
		return "Хорошист"
	}
	return "Троечник"
}

func main() {
	student := Student{
		Name:     "Вейнмайер Андрей",
		Age:      20,
		Course:   2,
		AvgGrade: 4.8,
	}

	fmt.Printf("%s: %s\n", student.Name, student.GetStatus())
}
