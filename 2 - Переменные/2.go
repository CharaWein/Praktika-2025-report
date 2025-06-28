package main

import (
	"fmt"
	"math"
)

func main() {
	var i float64 = 9999
	const pi = math.Pi
	const e = math.E

	fmt.Printf("π = %.5f, e = %.5f\n", pi, e)
	fmt.Printf("Площадь круга с радиусом %.0f: %.2f\n", i, pi*(math.Pow(i, 2)))
}
