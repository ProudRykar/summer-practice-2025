package main

import (
	"fmt"
	"math"
)

func main() {
	const pi = math.Pi
	const e = math.E

	radius := 5.0
	area := pi * radius * radius
	exp := math.Pow(e, 2)

	fmt.Println("Радиус круга:", radius)
	fmt.Printf("Площадь круга: %.2f\n", area)
	fmt.Printf("e^2: %.4f\n", exp)
}
