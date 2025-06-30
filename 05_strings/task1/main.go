package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "Hello, World!"

	// Подсчёт символов
	len := len(text)
	fmt.Println("Длина строки:", len)

	// Поиск подстроки
	str := strings.Contains(text, "World")
	fmt.Println("Содержит 'World':", str)

	// Изменение регистра
	fmt.Println("В верхнем регистре:", strings.ToUpper(text))
	fmt.Println("В нижнем регистре:", strings.ToLower(text))
}
