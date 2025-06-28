package main

import "fmt"

type Transport interface {
	Move()
	Stop()
}

type Car struct{}
type Bicycle struct{}

func (c Car) Move() {
	fmt.Println("Автомобиль поехал")
}

func (c Car) Stop() {
	fmt.Println("Автомобиль остановился")
}

func (b Bicycle) Move() {
	fmt.Println("Велосипед поехал")
}

func (b Bicycle) Stop() {
	fmt.Println("Велосипед остановился")
}

func main() {
	car := Car{}
	bicycle := Bicycle{}

	for i := 0; i < 1000000; i++ {
		// Четные итерации: автомобиль движется, велосипед останавливается
		if i%2 == 0 {
			car.Move()
			bicycle.Stop()
		} else {
			// Нечетные итерации: автомобиль останавливается, велосипед движется
			car.Stop()
			bicycle.Move()
		}
	}
}
