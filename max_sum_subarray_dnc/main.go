/* 53. Maximum Subarray

 */

package main

import (
	"fmt"
	"math"
)

func maxSubArray(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	midpoint := len(nums) / 2
	left := nums[0:midpoint] // [-1]
	right := nums[midpoint:] //[0]

	leftResult := maxSubArray(left)   // [-2], -2
	rightResult := maxSubArray(right) // [-1, 0]

	return combine(leftResult, rightResult, nums)
}

func combine(ls, rs int, nums []int) int {
	midpoint := len(nums)/2 - 1

	maxLeftSum := math.MinInt32
	leftCurrentVal := 0

	for i := midpoint; i >= 0; i-- {
		leftCurrentVal += nums[i]

		if leftCurrentVal > maxLeftSum {
			maxLeftSum = leftCurrentVal
		}
	}

	maxRightSum := math.MinInt32
	rightCurrentVal := 0

	for i := midpoint + 1; i < len(nums); i++ {
		rightCurrentVal += nums[i]

		if rightCurrentVal > maxRightSum {
			maxRightSum = rightCurrentVal
		}
	}

	cs := maxLeftSum + maxRightSum

	if ls >= cs && ls >= rs {
		return ls
	} else if cs >= ls && cs >= rs {
		return cs
	} else {
		return rs
	}
}

func main() {
	fmt.Println(maxSubArray([]int{-2, -5, 6, -2, -3, 1, 5, -6}))   // 7
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})) // 6
	fmt.Println(maxSubArray([]int{-2, 1, 0}))                      // 1
}
