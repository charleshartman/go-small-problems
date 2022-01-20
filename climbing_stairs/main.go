package main

import (
	"fmt"

	"github.com/charleshartman/tortoise"
)

// Skeleton to test importing custom library

func main() {
	fmt.Println("Hello world")
	fmt.Println(tortoise.FactorialInt(5))
	fmt.Println(tortoise.MaxInt(51, 132))
	fmt.Println(tortoise.MinInt(51, 132))
}
