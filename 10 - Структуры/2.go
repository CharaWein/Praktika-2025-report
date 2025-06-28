package main

import "fmt"

type Engine struct {
	Power int
	Type  string
}

type Car struct {
	Model  string
	Year   int
	Engine Engine
}

func main() {
	car := Car{
		Model: "Москвич-412",
		Year:  1966,
		Engine: Engine{
			Power: 75,
			Type:  "УЗАМ-412",
		},
	}

	fmt.Printf("%s %d, двигатель: %s %d л.с.\n",
		car.Model, car.Year, car.Engine.Type, car.Engine.Power)
}
