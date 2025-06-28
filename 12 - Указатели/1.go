package main

import "fmt"

func swap(a, b *int) {
	*a, *b = *b, *a
}

func main() {
	x, y := 19, 92
	fmt.Println("До:", x, y)
	swap(&x, &y)
	fmt.Println("После:", x, y)
}
