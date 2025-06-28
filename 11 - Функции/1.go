package main

import (
	"fmt"
	"sort"
)

// Функция поиска элемента в срезе
func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

// Функция фильтрации среза по условию
func filter(slice []string, condition func(string) bool) []string {
	var result []string
	for _, s := range slice {
		if condition(s) {
			result = append(result, s)
		}
	}
	return result
}

// Функция сортировки среза
func sortStrings(slice []string) []string {
	sorted := make([]string, len(slice))
	copy(sorted, slice)
	sort.Strings(sorted)
	return sorted
}

func main() {
	langs := []string{"Андрей", "И", "Чapa", "Друзья"}

	// Поиск
	fmt.Println("Содержит 'Вейнмайер':", contains(langs, "Вейнмайер"))

	// Сортировка
	sorted := sortStrings(langs)
	fmt.Println("Отсортировано:", sorted)

	// Фильтрация
	shortLangs := filter(langs, func(s string) bool {
		return len(s) <= 3
	})
	fmt.Println("Короткие слова:", shortLangs)
}
