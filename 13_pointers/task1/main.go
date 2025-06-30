package main

import "fmt"

// ФУНКЦИЯ: обмен значениями двух переменных по указателям
func swap(a, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

func main() {
	x, y := 10, 20
	fmt.Printf("До обмена: x = %d, y = %d\n", x, y)
	swap(&x, &y)
	fmt.Printf("После обмена: x = %d, y = %d\n", x, y)
}
