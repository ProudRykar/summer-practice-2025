package main

import (
	"fmt"
)

// ФУНКЦИЯ: добавлениe оценки
func addGrade(grades map[string]int, name string, grade int) {
	grades[name] = grade
}

// ФУНКЦИЯ: удалениe студента
func removeStudent(grades map[string]int, name string) {
	delete(grades, name)
}

// ФУНКЦИЯ: поиск оценки по имени
func getGrade(grades map[string]int, name string) {
	grade, exists := grades[name]
	if exists {
		fmt.Printf("Оценка для %s: %d\n", name, grade)
	} else {
		fmt.Printf("Студент %s не найден\n", name)
	}
}

func main() {
	grades := make(map[string]int)

	addGrade(grades, "Олег", 69)
	addGrade(grades, "Степан", 14)

	getGrade(grades, "Олег")
	getGrade(grades, "Игорь") // ОШИБКА: Нет такого студента

	removeStudent(grades, "Степан")

	getGrade(grades, "Степан") // ОШИБКА: Студент уже был удалён
}
