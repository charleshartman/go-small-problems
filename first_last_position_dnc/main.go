/* 34. Find First and Last Position of Element in Sorted Array

UNDERSTAND THE PROBLEM
  Inputs
    - slice of ints, target integer
  Outputs
    - slice of two integers, representing the start and end indexes of the target integer within the input slice

Rules/Implicit Requirements
  Explicit
    - nums is a non-decreasing array
    - return a slice of [-1, -1] if the target is not found
    - 0 <= nums.length <= 105
    - -109 <= nums[i] <= 109
    - -109 <= target <= 109
  Implicit
    - an empty slice returns a slice of [-1, -1] irrespective of the target
    - if the target element appears only once, the same index should be returned as both start and end indexes

Mental Model (optional)
  - main function calls two helper functions that are invoked recursively, bubbling up index values that main function returns as result (two element), helper functions receive entire input slice, pointers for index values adjust with each recursive call to find left most and right most target value depending on which helper function we are utilizing
    - helper function i/o
      - input: slice of ints, target integer, left pointer integer, right pointer integer
      - output: single integer (each), left and right index

  - Each recursive call adjusts the left and right pointers based on the following conditions:
    - To find the leftmost target: if you find the target or find an element *greater than* the target, move to the left. Otherwise move to the right.
    - To find the rightmost target: if you find the target or find an element *less than* or equal to the target, move to the right. Otherwise move to the left.
    - Base Case: If the left pointer is greater or equal to the right pointer
      - If the value at the left pointer is equal to the target, return the pointer
      - Else, the value is not present in the slice

EXAMPLES/TEST CASES
  Expected Cases
  - nums = []int{5,7,7,8,8,10}
  - target = 8
  - return value = []int{3, 4}

  - nums = []int{5,7,7,8,8,10}
  - target = 6
  - return value = []int{-1, -1}

  - nums = []int{8,8}
  - target = 8
  - return value = []int{0, 1}

  Edge Cases
  - nums = []int{}
  - target = 0
  - return value = []int{-1, -1}

  - nums = []int{8}
  - target = 8
  - return value = []int{0, 0}

ALGORITHM
  - guard clause for input slices of length 0
    - return a slice of [-1, -1]

    - helper function (findStart)
      - find the index of the leftmost target
        - pass in starting and ending indexes for left and right pointers
        - base case: if the left pointer and right pointer meet/cross
          - evaluate whether the value at the left pointer == target
            - if it is, we've found the leftmost target => return the leftP
            - if it isn't, the target doesn't exist => return -1
        - calculate the midpoint based on the current values of the left and right pointers
        - if the value at the midpoint is less than the target
          - move the left pointer to the midpoint + 1
        - else (midpoint is >= target)
          - move the right pointer to the midpoint
        - call the helper function again with the new left and right pointers

    - helper function (findEnd)
      - find the index of the rightmost target
      - pass in starting and ending indexes for left and right pointers
        - base case: if the left pointer and right pointer meet/cross
          - evaluate whether the value at the left pointer == target
            - if it is, we've found the rightmost target => return the leftP
            - if it isn't, the target doesn't exist => return -1
        - calculate the midpoint based on the current values of the left and right pointers
        - if the value at the midpoint is greater than the target
          - move the right pointer to the midpoint - 1
        - else (midpoint is <= target)
          - move the left pointer to the midpoint
        - call the helper function again with the new left and right pointers

  - return the indexes of the leftmost and rightmost target in a slice

CODE
*/

package main

import "fmt"

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	start := findStart(nums, target, 0, len(nums)-1)
	end := findEnd(nums, target, 0, len(nums)-1)

	return []int{start, end}
}

func findStart(nums []int, target, leftP, rightP int) int {
	if leftP >= rightP {
		if nums[leftP] == target {
			return leftP
		} else {
			return -1
		}
	}

	midpoint := (leftP + rightP) / 2

	if nums[midpoint] < target {
		leftP = midpoint + 1
	} else {
		rightP = midpoint
	}

	return findStart(nums, target, leftP, rightP)
}

func findEnd(nums []int, target, leftP, rightP int) int {
	if leftP >= rightP {
		if nums[leftP] == target {
			return leftP
		} else {
			return -1
		}
	}

	midpoint := (leftP+rightP)/2 + 1

	if nums[midpoint] > target {
		rightP = midpoint - 1
	} else {
		leftP = midpoint
	}

	return findEnd(nums, target, leftP, rightP)
}

func main() {
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 8)) // [3, 4]
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 6)) // [-1, -1]
	fmt.Println(searchRange([]int{8, 8}, 8))              // [0, 1]
	fmt.Println(searchRange([]int{}, 0))                  // [-1, -1]
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 5)) // [0, 0]
	fmt.Println(searchRange([]int{1}, 0))                 // [-1, -1]
	fmt.Println(searchRange([]int{1}, 1))                 // [0, 0]
	fmt.Println(searchRange([]int{1, 4}, 4))              // [1, 1]
}
