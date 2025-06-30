package main

import (
	"fmt"
	"sort"
)

// ФУНКЦИЯ: поиск элемента в срезе строк
	// RETURNS: индекс или -1 если не найден
func findElement(slice []string, target string) int {
	for i, v := range slice {
		if v == target {
			return i
		}
	}
	return -1
}

// ФУНКЦИЯ: сортировка среза строк по алфавиту
func sortSlice(slice []string) []string {
	sorted := make([]string, len(slice))
	copy(sorted, slice)
	sort.Strings(sorted)
	return sorted
}

// ФУНКЦИЯ: фильтрация среза
//		RETURNS: срез, содержащий элементы, длина которых >= minLen
func filterByLength(slice []string, minLen int) []string {
	var filtered []string
	for _, v := range slice {
		if len(v) >= minLen {
			filtered = append(filtered, v)
		}
	}
	return filtered
}

func main() {
	data := []string{"apple", "banana", "pear", "kiwi", "orange"}

	fmt.Println("Исходный срез:", data)

	pos := findElement(data, "pear")
	fmt.Println("Индекс элемента 'pear':", pos)

	sorted := sortSlice(data)
	fmt.Println("Отсортированный срез:", sorted)

	filtered := filterByLength(data, 5)
	fmt.Println("Отфильтрованный срез (длина >= 5):", filtered)
}
