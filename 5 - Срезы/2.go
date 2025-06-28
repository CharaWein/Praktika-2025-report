package main

import "fmt"

func remove(slice []string, index int) []string {
	return append(slice[:index], slice[index+1:]...)
}

func main() {
	slice := []string{"Андрей", "и", "Чара", "не", "друзья"}
	slice = remove(slice, 3)

	fmt.Println(slice)
}
