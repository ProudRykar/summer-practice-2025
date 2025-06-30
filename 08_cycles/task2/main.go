package main

import "fmt"

func main() {
	const limit = 100

	for num := 2; num <= limit; num++ {
		isPrime := true
		for i := 2; i*i <= num; i++ {
			if num%i == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			fmt.Println(num)
		}
	}
}
