package main

import (
	"fmt"
)

// ФУНКЦИЯ: удаляет элемент по индексу из среза строк
	// RETURNS: новый срез без элемента по индексу
func removeAtIndex(slice []string, index int) []string {
	if index < 0 || index >= len(slice) {
		fmt.Println("Недопустимый индекс")
		return slice
	}

	return append(slice[:index], slice[index+1:]...)
}

func main() {
	slice := []string{"Go", "is", "very", "fast"}

	fmt.Println("До удаления:", slice)

	slice = removeAtIndex(slice, 2) // Удалим "very"

	fmt.Println("После удаления:", slice)
}
