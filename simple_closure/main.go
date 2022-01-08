package main

import "fmt"

func returnFunc() func() {
	var sum int = 0

	return func() {
		sum++
		fmt.Println(sum)
	}
}

func main() {
	x := returnFunc()
	x() // 1
	x() // 2
	x() // 3
}
