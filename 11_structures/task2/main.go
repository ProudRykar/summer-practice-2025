package main

import "fmt"

// СТРУКТУРА: Двигатель
type Engine struct {
	Horsepower int
	Type       string
}

// СТРУКТУРА: Машина
type Car struct {
	Make   string
	Model  string
	Year   int
	Engine Engine
}

func main() {
	car := Car{
		Make:  "Toyota",
		Model: "Camry",
		Year:  2021,
		Engine: Engine{
			Horsepower: 203,
			Type:       "Hybrid",
		},
	}

	fmt.Printf("Автомобиль: %s %s (%d)\n", car.Make, car.Model, car.Year)
	fmt.Printf("Двигатель: %d л.с., тип: %s\n", car.Engine.Horsepower, car.Engine.Type)
}
