package main

import (
	"fmt"
	"strings"
)

// ФУНКЦИЯ: подсчёт слов в строке
func countWords(text string) map[string]int {
	wordFreq := make(map[string]int)

	words := strings.Fields(strings.ToLower(text))

	for _, word := range words {
		wordFreq[word]++
	}

	return wordFreq
}

func main() {
	text := "Go is fun and go is fast and fun"

	wordFreq := countWords(text)

	fmt.Println("Частота слов:")
	for word, count := range wordFreq {
		fmt.Printf("%s: %d\n", word, count)
	}
}
