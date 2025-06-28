package main

import "fmt"

func byValue(n int) {
	n = 0
}

func byPointer(n *int) {
	*n = 19
}

func main() {
	num := 9999

	byValue(num)
	fmt.Println("По значению:", num)

	byPointer(&num)
	fmt.Println("По указателю:", num)
}
