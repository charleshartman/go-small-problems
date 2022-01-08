package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(DigPow(89, 1))    // should return 1 since 8¹ + 9² = 89 = 89 * 1
	fmt.Println(DigPow(92, 1))    // should return -1 since there is no k such as 9¹ + 2² equals 92 * k
	fmt.Println(DigPow(695, 2))   // should return 2 since 6² + 9³ + 5⁴= 1390 = 695 * 2
	fmt.Println(DigPow(46288, 3)) // should return 51 since 4³ + 6⁴+ 2⁵ + 8⁶ + 8⁷ = 2360688 = 46288 * 51
}

func DigPow(n, p int) int {
	numStrs := strings.Split(strconv.Itoa(n), "")
	nums := make([]float64, len(numStrs))
	for i, v := range numStrs {
		num, _ := strconv.ParseFloat(v, 64)
		nums[i] = num
	}

	var sum float64
	exp := float64(p)
	for _, v := range nums {
		sum += math.Pow(v, exp)
		exp++
	}
	sumInt := int(sum)

	if sumInt%n == 0 {
		return sumInt / n
	}

	return -1
}

/* algo
- convert num to string, split string, convert strings to ints
- declare var sum int
- loop through digits
	- mult by mult and add to sum
	- add one to mult
- now with sum:
	- count = 1
	- running total = 0
	- while running total <= sum
	- if sum is equal to num * count
		return count
	- if no loop returns then return -1
*/
