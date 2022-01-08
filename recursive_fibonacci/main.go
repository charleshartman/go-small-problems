package main

import "fmt"

func main() {
	memo := make(map[int]int)
	fmt.Println(fibonacci(20, memo)) // 6765
}

func fibonacci(n int, m map[int]int) int {
	if n == 0 || n == 1 {
		// add to memo just so hash is complete because... must have them all
		m[n] = n
		return n
	}

	_, ok := m[n]
	if !ok {
		m[n] = fibonacci(n-1, m) + fibonacci(n-2, m)
	}

	// to log all the beautiful fibs
	fmt.Println(m)

	return m[n]
}
