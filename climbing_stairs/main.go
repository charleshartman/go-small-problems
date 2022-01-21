package main

import (
	"fmt"

	"github.com/charleshartman/tortoise"
)

func climbingStairs(n int) int {
	dp := map[int]int{}
	dp[1], dp[2] = 1, 2

	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	fmt.Println(dp)
	return dp[n]
}

// Using this problem package to test importing custom library as well

func main() {
	fmt.Println("Hello world")
	fmt.Println(tortoise.FactorialInt(5))
	fmt.Println(tortoise.MaxInt(51, 132))
	fmt.Println(tortoise.MinInt(51, 132))
	fmt.Println(climbingStairs(7))
}
