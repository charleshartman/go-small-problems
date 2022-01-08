package main

import "fmt"

func CountPositivesSumNegatives(numbers []int) []int {
	var res []int
	var posCnt, negSum int
	for _, val := range numbers {
		if val > 0 {
			posCnt += 1
		} else {
			negSum += val
		}
	}
	res = append(res, posCnt, negSum)
	return res
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, -11, -12, -13, -14, -15}
	fmt.Println(CountPositivesSumNegatives(nums))
}
