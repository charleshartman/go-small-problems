package main

import "fmt"

func NbYear(p0 int, percent float64, aug int, p int) int {
	percent = percent * .01
	var calcP0, calcAug, calcP float64 = float64(p0), float64(aug), float64(p)
	var years int
	for calcP0 < calcP {
		calcP0 += (calcP0 * percent) + calcAug
		years++
	}
	return years
}

func main() {
	fmt.Println(NbYear(1500, 5, 100, 5000) == 15)
	fmt.Println(NbYear(1500000, 2.5, 10000, 2000000) == 10)
	fmt.Println(NbYear(1500000, 0.25, 1000, 2000000) == 94)
	fmt.Println(NbYear(1500000, 0.25, -1000, 2000000) == 151)
}
