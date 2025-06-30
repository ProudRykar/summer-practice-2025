package main

import (
	"fmt"
)

func main() {
	var a, b float64
	var op string

	fmt.Print("Введите первое число: ")
	fmt.Scan(&a)

	fmt.Print("Введите оператор (+, -, *, /): ")
	fmt.Scan(&op)

	fmt.Print("Введите второе число: ")
	fmt.Scan(&b)

	switch op {
	case "+":
		fmt.Printf("%.2f + %.2f = %.2f\n", a, b, a+b)
	case "-":
		fmt.Printf("%.2f - %.2f = %.2f\n", a, b, a-b)
	case "*":
		fmt.Printf("%.2f * %.2f = %.2f\n", a, b, a*b)
	case "/":
		if b != 0 {
			fmt.Printf("%.2f / %.2f = %.2f\n", a, b, a/b)
		} else {
			fmt.Println("Ошибка: деление на ноль!")
		}
	default:
		fmt.Println("Неверный оператор")
	}
}
