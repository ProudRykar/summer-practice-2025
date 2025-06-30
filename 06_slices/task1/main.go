package main

import (
	"fmt"
)

func main() {
	// Пустой срез
	slice := []string{}

	// Динамически добавляем элементы
	slice = append(slice, "Go")
	slice = append(slice, "is")
	slice = append(slice, "awesome!")

	// Выводим все элементы среза
	fmt.Println("Содержимое среза:")
	for i, v := range slice {
		fmt.Printf("[%d]: %s\n", i, v)
	}
}
