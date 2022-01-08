package main

import "fmt"

func Tribonacci(signature [3]float64, n int) []float64 {
	seq := []float64{signature[0], signature[1], signature[2]}

	for len(seq) < n {
		length := len(seq)
		seq = append(seq, seq[length-1]+seq[length-2]+seq[length-3])
	}

	return seq[:n]
}

func main() {
	fmt.Println(Tribonacci([3]float64{1, 1, 1}, 10))
	fmt.Println(Tribonacci([3]float64{0, 0, 1}, 10))
	fmt.Println(Tribonacci([3]float64{0, 0, 0}, 10))
	fmt.Println(Tribonacci([3]float64{1, 1, 1}, 1))
	fmt.Println(Tribonacci([3]float64{1, 1, 1}, 2))
	fmt.Println(Tribonacci([3]float64{0.5, 0.5, 0.5}, 30))
}
