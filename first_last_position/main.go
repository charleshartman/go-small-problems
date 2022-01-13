/* 34. Find First and Last Position of Element in Sorted Array

Given an array of integers nums sorted in non-decreasing order, find the
starting and ending position of a given target value.

If target is not found in the array, return [-1, -1].

You must write an algorithm with O(log n) runtime complexity.

Example 1:
Input: nums = [5,7,7,8,8,10], target = 8
Output: [3,4]

Example 2:
Input: nums = [5,7,7,8,8,10], target = 6
Output: [-1,-1]

Example 3:
Input: nums = [], target = 0
Output: [-1,-1]

Example 4
first, last
[5,7,7,8,8,8,8,8,10], target = 8
 [5,7,7,8] 8 [8,8,8,10]
 [5] 7 [7,8] 8 [8,8,8,10]
        7[8] 8 [8,8,8,10]
             8 [8] 8 [8, 10]
             8 [8] 8 8 [10]




 Constraints:
0 <= nums.length <= 105
-109 <= nums[i] <= 109
nums is a non-decreasing array.
-109 <= target <= 109

ALGORITHM
  -
Find mid point
- binary search until target is found

- if the nums[midPointIndex - 1] == target
  - need to binary search left
  - else first = midPointIndex

- if the nums[midPointIndex + 1] == target
  - need to binary search right
  - else last = midPointIndex
CODE

*/

package main

import "fmt"

func searchRange(nums []int, target int) []int {
	if len(nums) == 1 {
		return []int{0, 0}
	}

	if target == 0 {
		return []int{-1, -1}
	}

	start := binarySearchLeft(nums, target)
	end := binarySearchRight(nums, target)

	return []int{start, end}
}

func binarySearchLeft(slice []int, target int) int {
	left, right := 0, len(slice)-1

	for left <= right {
		mid := left + ((right - left) / 2)

		if mid == 0 {
			break
		}

		if slice[mid] == target && slice[mid-1] != target {
			return mid
		} else if target > slice[mid] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}

func binarySearchRight(slice []int, target int) int {
	left, right := 0, len(slice)-1

	for left <= right {
		mid := left + ((right - left) / 2)

		if mid >= len(slice)-1 {
			break
		}

		if slice[mid] == target && slice[mid+1] != target {
			return mid
		} else if target > slice[mid] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}

func main() {
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 8))    // [3,4]
	fmt.Println(searchRange([]int{5, 7, 7, 7, 8, 8, 10}, 7)) // [1,3]
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 6))    // [-1,-1]
	fmt.Println(searchRange([]int{}, 0))                     // [-1,-1]
	// fmt.Println(searchRange([]int{2, 2}, 2))                 // [0,1]
	// fmt.Println(searchRange([]int{9, 9, 9, 9, 9}, 9))        // [0,4]
}
