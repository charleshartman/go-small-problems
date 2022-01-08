package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Min(s []int) int {
	minVal := s[0]
	for _, v := range s {
		if v < minVal {
			minVal = v
		}
	}
	return minVal
}

func Max(s []int) int {
	maxVal := s[0]
	for _, v := range s {
		if v > maxVal {
			maxVal = v
		}
	}
	return maxVal
}

func HighAndLow(str string) string {
	numStrings := strings.Split(str, " ")
	numInts := []int{}
	for _, v := range numStrings {
		v, _ := strconv.Atoi(v)
		numInts = append(numInts, v)
	}
	return strconv.Itoa(Max(numInts)) + " " + strconv.Itoa(Min(numInts))
}

func main() {
	fmt.Println(HighAndLow("1 2 3 4 5") == "5 1")
	fmt.Println(HighAndLow("1 2 -3 4 5") == "5 -3")
	fmt.Println(HighAndLow("1 9 3 4 -5") == "9 -5")
}
