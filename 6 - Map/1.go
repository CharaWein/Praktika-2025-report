package main

import "fmt"

func main() {
	grades := make(map[string]int)

	// Добавление
	grades["Андрей"] = 5
	grades["Чара"] = 5
	grades["Оля"] = 4
	grades["Максим"] = 3
	grades["Ульяна"] = 2
	grades["Наташа"] = 1

	// Поиск
	if grade, exists := grades["Чара"]; exists {
		fmt.Println("Оценка Чары:", grade)
	}

	// Удаление
	delete(grades, "Наташа")

	fmt.Println(grades)
}
