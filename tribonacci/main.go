package main

import "fmt"

func Tribonacci(signature [3]float64, n int) []float64 {
	seq := signature[:]

	for len(seq) < n {
		length := len(seq)
		seq = append(seq, seq[length-1]+seq[length-2]+seq[length-3])
	}

	return seq[:n]
}

func main() {
	fmt.Println(Tribonacci([3]float64{1, 1, 1}, 10))       // [1 1 1 3 5 9 17 31 57 105]
	fmt.Println(Tribonacci([3]float64{0, 0, 1}, 10))       // [0 0 1 1 2 4 7 13 24 44]
	fmt.Println(Tribonacci([3]float64{0, 0, 0}, 10))       // [0 0 0 0 0 0 0 0 0 0]
	fmt.Println(Tribonacci([3]float64{1, 1, 1}, 1))        // [1]
	fmt.Println(Tribonacci([3]float64{1, 1, 1}, 2))        // [1 1]
	fmt.Println(Tribonacci([3]float64{0.5, 0.5, 0.5}, 30)) // [0.5 0.5 0.5 1.5 2.5 4.5 8.5 15.5 28.5 52.5 96.5 177.5 326.5 600.5 1104.5 2031.5 3736.5 6872.5 12640.5 23249.5 42762.5 78652.5 144664.5 266079.5 489396.5 900140.5 1.6556165e+06 3.0451535e+06 5.6009105e+06 1.03016805e+07]
}
