package main

import "fmt"

func main() {
	a, b := 9, -9

	fmt.Printf("%d + %d = %d\n", a, b, a+b)
	fmt.Printf("%d - %d = %d\n", a, b, a-b)
	fmt.Printf("%d * %d = %d\n", a, b, a*b)
	fmt.Printf("%d / %d = %d\n", a, b, a/b)
	fmt.Printf("%d %% %d = %d\n", a, b, a%b)
	fmt.Printf("%d & %d = %d\n", a, b, a&b)
	fmt.Printf("%d | %d = %d\n", a, b, a|b)
	fmt.Printf("%d ^ %d = %d\n", a, b, a^b)
}
