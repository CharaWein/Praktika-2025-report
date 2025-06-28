package main

import "fmt"

func main() {
	var i int = 9999
	var f float64 = 9999.0
	var b bool = true
	var s string = "Чара"
	var r rune = 'В'
	var by byte = 255

	fmt.Printf("int: %d, float64: %f, bool: %t, string: %s, rune: %c, byte: %d\n",
		i, f, b, s, r, by)
}
