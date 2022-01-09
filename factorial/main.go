package main

import "fmt"

func Factorial(n int) int {
	product := 1
	for i := 2; i <= n; i++ {
		product *= i
	}

	return product
}

func main() {
	fmt.Println(Factorial(10)) // 3628800
}
