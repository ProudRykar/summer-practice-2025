package main

import (
	"fmt"
)

// ФУНКЦИЯ: Нахождение самой длинной строки
func longestString(strings ...string) string {
	longest := ""
	for _, s := range strings {
		if len(s) > len(longest) {
			longest = s
		}
	}
	return longest
}

func main() {
	result := longestString("cat", "elephant", "dog", "hippopotamus", "mouse")
	fmt.Println("Самая длинная строка:", result)
}
