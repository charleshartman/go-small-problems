package main

import "fmt"

func FindOdd(seq []int) int {
	var oddCount int
	hash := make(map[int]int)
	for _, v := range seq {
		_, ok := hash[v]
		if ok {
			hash[v] += 1
		} else {
			hash[v] = 1
		}
	}

	for k := range hash {
		if hash[k]%2 == 1 {
			oddCount = k
			break
		}
	}
	return oddCount
}

func main() {
	seq1 := []int{20, 1, -1, 2, -2, 3, 3, 5, 5, 1, 2, 4, 20, 4, -1, -2, 5}
	seq2 := []int{1, 1, 2, -2, 5, 2, 4, 4, -1, -2, 5}
	seq3 := []int{20, 1, 1, 2, 2, 3, 3, 5, 5, 4, 20, 4, 5}
	seq4 := []int{10}
	seq5 := []int{1, 1, 1, 1, 1, 1, 10, 1, 1, 1, 1}
	seq6 := []int{5, 4, 3, 2, 1, 5, 4, 3, 2, 10, 10}
	fmt.Println(FindOdd(seq1))
	fmt.Println(FindOdd(seq2))
	fmt.Println(FindOdd(seq3))
	fmt.Println(FindOdd(seq4))
	fmt.Println(FindOdd(seq5))
	fmt.Println(FindOdd(seq6))
}
