package main

import "fmt"

func Factorial(n float64) float64 {
	var product float64 = 1
	var i float64
	for i = 2; i <= n; i++ {
		product *= i
	}

	return product
}

func Combinations(elements, choices float64) float64 {
	//  C(n,r) = n! / r! (n – r)!
	return Factorial(elements) / (Factorial(choices) * Factorial(elements-choices))
}

func main() {
	fmt.Println(Combinations(5, 3) == 10)
	fmt.Println(Combinations(4, 2) == 6)
	fmt.Println(Combinations(4, 3) == 4)
	fmt.Println(Combinations(4, 0) == 1)
	fmt.Println(Combinations(7, 2) == 21)
	fmt.Println(Combinations(6, 2) == 15)
	// int data type would fine fine up until this point
	// but the following example requires larger number, so we use float64
	fmt.Println(Combinations(25, 3) == 2300)
}
