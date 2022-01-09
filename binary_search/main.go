package main

import (
	"errors"
	"fmt"
)

func ValidateSliceIsSorted(sl []int) bool {
	for i := 0; i < len(sl)-1; i++ {
		if sl[i] > sl[i+1] {
			return false
		}
	}
	return true
}

func BinarySearch(sl []int, n int) (int, error) {
	if !ValidateSliceIsSorted(sl) {
		return -1, errors.New("Argument is not sorted, BinarySearch only works for sorted datasets, exiting...")
	}

	lowerBound := 0
	upperBound := len(sl) - 1

	for lowerBound <= upperBound {
		midpoint := (upperBound + lowerBound) / 2

		valueAtMidpoint := sl[midpoint]

		if n == valueAtMidpoint {
			return midpoint, nil
		} else if n < valueAtMidpoint {
			upperBound = midpoint - 1
		} else if n > valueAtMidpoint {
			lowerBound = midpoint + 1
		}
	}

	return -1, errors.New("Argument was not found in the slice.")
}

func main() {
	dataset1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	dataset2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	dataset3 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 13, 17, 21, 35, 101}
	dataset4 := []int{1, 3, 4, 5, 6, 7, 8, 9, 11, 15}
	dataset5 := []int{1, 2, 3, 4, 6, 5, 7, 8, 9}

	fmt.Println(BinarySearch(dataset1, 7))
	fmt.Println(BinarySearch(dataset2, 11))
	fmt.Println(BinarySearch(dataset3, 2))
	fmt.Println(BinarySearch(dataset4, 6))
	fmt.Println(BinarySearch(dataset4, 15))
	fmt.Println(BinarySearch(dataset5, 6))
}
