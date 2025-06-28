package main

import "fmt"

func main() {
	var slice []string

	slice = append(slice, "Андрей")
	slice = append(slice, "друг")
	slice = append(slice, "Чары")

	for i, item := range slice {
		fmt.Printf("%d: %s\n", i, item)
	}
}
