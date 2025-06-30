package main

import (
	"fmt"
)

func main() {
	a := 15
	b := 4

	fmt.Printf("%d + %d = %d\n", a, b, a+b)
	fmt.Printf("%d - %d = %d\n", a, b, a-b)
	fmt.Printf("%d * %d = %d\n", a, b, a*b)
	fmt.Printf("%d / %d = %d\n", a, b, a/b)
	fmt.Printf("%d %% %d = %d\n", a, b, a%b)

	a++
	fmt.Println("a после a++:", a)
	b--
	fmt.Println("b после b--:", b)
}
