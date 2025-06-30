package main

import (
	"fmt"
	"math"
)

// ИНТЕРФЕЙС: Форма
type Shape interface {
	Area() float64
}

// СТРУКТУРА: прямоугольник
type Rectangle struct {
	Width, Height float64
}

// ФУНКЦИЯ: вычисление площади прямоугольника
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// СТРУКТУРА: круг
type Circle struct {
	Radius float64
}

// ФУНКЦИЯ: вычисление площади круга
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func printArea(s Shape) {
	fmt.Printf("Площадь: %.2f\n", s.Area())
}

func main() {
	rect := Rectangle{Width: 5, Height: 3}
	circ := Circle{Radius: 4}

	printArea(rect)
	printArea(circ)
}
