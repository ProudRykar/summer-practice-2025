package main

import (
	"fmt"
	"time"
)

// СТРУКТУРА: Студент
type Student struct {
	Name      string
	BirthYear int
	AvgGrade  float64
}

// ФУНКЦИЯ: вычисление текущего возраста студента
func (s Student) Age() int {
	currentYear := time.Now().Year()
	return currentYear - s.BirthYear
}

// ФУНКЦИЯ: получение статуса по среднему баллу
func (s Student) Status() string {
	switch {
	case s.AvgGrade >= 4.5:
		return "Отличник"
	case s.AvgGrade >= 3.5:
		return "Хорошист"
	default:
		return "Троечник" // ну то есть - я
	}
}

func main() {
	student := Student{Name: "Никита", BirthYear: 2005, AvgGrade: 3.9}

	fmt.Printf("Студент: %s\n", student.Name)
	fmt.Printf("Возраст: %d\n", student.Age())
	fmt.Printf("Статус: %s\n", student.Status())
}
