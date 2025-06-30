package main

import (
	"fmt"
	"strings"
)

func main() {
	sentence := "MTUСI is Russia's leading sectoral university in the field of telecommunications, information technologies and information security."

	// Разбиваем строку на слова
	words := strings.Split(sentence, " ")

	// Выводим каждое слово на новой строке
	fmt.Println("Слова в предложении:")
	for _, word := range words {
		fmt.Println(word)
	}
}
