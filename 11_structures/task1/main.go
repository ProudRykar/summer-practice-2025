package main

import "fmt"

// СТРУКТУРА: Студент
type Student struct {
	Name     string
	Age      int
	Course   int
	AvgGrade float64
}

// ФУНКЦИЯ: вывод информации о студенте
func printStudent(s Student) {
	fmt.Printf("Имя: %s\nВозраст: %d\nКурс: %d\nСредний балл: %.2f\n",
		s.Name, s.Age, s.Course, s.AvgGrade)
}

// ФУНКЦИЯ: обновление среднего балла
func updateAvgGrade(s *Student, newGrade float64) {
	s.AvgGrade = newGrade
}

func main() {
	student := Student{
		Name:     "Антон",
		Age:      20,
		Course:   2,
		AvgGrade: 4.5, // точно не соврал)
	}

	printStudent(student)

	updateAvgGrade(&student, 4.8)

	fmt.Println("После обновления среднего балла:")
	printStudent(student)
}
