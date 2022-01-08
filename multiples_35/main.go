package main

import "fmt"

func Multiple3And5(number int) int {
	resultSum := 0
	for i := 1; i < number; i++ {
		if i%3 == 0 || i%5 == 0 {
			resultSum += i
		}
	}
	return resultSum
}

func main() {
	fmt.Println(Multiple3And5(10)) // 23
	fmt.Println(Multiple3And5(15)) // 45
	fmt.Println(Multiple3And5(32)) // 225
}
