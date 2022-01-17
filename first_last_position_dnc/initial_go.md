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

Constraints:
0 <= nums.length <= 105
-109 <= nums[i] <= 109
nums is a non-decreasing array.
-109 <= target <= 109

Example Breakdown
        F F F T T F        midpoint = len / 2 = index 2
       [5,7,7,8,8,10] --> (3 , 4)
  [5,7,7]-->(-1, -1)       [8,8,10] --> (0, 1)
[5]-->(-1, -1) [7, 7]-->(-1, -1)     [8]-->(0, 0) [8, 10] --> (0, 0)
      [7]-->(-1, -1) [7] -->(-1, -1)        [8]--> (0, 0) [10] -->(-1,-1)

            0   1  2   3  midpoint = 1
            [1, 8, 8, 10]

         0  1        2   3
         0  1        0   1
        [1, 8]      [8, 10]

        0     1    2      3
        0     0    0      0
      [1]    [8]  [8]    [10]       ---> Base Case


slice = [8, 10] = length of 2
someFunc([8, 10]) = 0, 0
return 2, 2


func searchRange(nums []int, target int) []int {
    if len(nums) == 0 {
        return []int{-1, -1}
    }
    return []int{binaryLeft(nums, target), binaryRight(nums, target)}
}

func binaryLeft(nums []int, target int) int {
    left, right := 0, len(nums)-1

    for left < right {
        mid := (right + left)/2
        if nums[mid] < target {
            left = mid + 1
        } else {
            right = mid
        }
    }
    if nums[left] == target {
        return left
    }
    return -1
}

func binaryRight(nums []int, target int) int {
    left, right := 0, len(nums)-1

    for left < right {
        mid := (right + left)/2 + 1
        if nums[mid] > target {
            right = mid - 1
        } else if nums[mid] <= target {
            left = mid
        }
    }
    if nums[left] == target {
        return left
    }
    return -1
}

- To find the leftmost target: if you find the target or find an element *greater than* the target, move to the left (i.e., right pointer = midpoint - 1)
- To find the rightmost target: if you find the target or find an element *less than* or equal to the target, move to the right (i.e., left pointer = midpoint + 1)

func dncProblem(list []int) []int {
  _// base cases; very different for every problem_

  _// divide_
  left := leftHalf(list)
  right := rightHalf(list)

  _// conquer (recurse)_
  leftResult := dncProblem(left)
  rightResult := dncProblem(right)

  _// combine; very different for every problem_
  return combine(leftResult, rightResult)
}

*/
```go
package main

import "fmt"

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 || len(nums) == 1 && nums[0] != target {
		return []int{-1, -1}
	}

	if len(nums) == 1 && nums[0] == target {
		return []int{0, 0}
	}

	leftIndex := leftSearch(nums, target, 0, len(nums)-1)
	rightIndex := rightSearch(nums, target, 0, len(nums)-1)

	return []int{leftIndex, rightIndex}
}

func leftSearch(nums []int, target, leftP, rightP int) int {
	mid := leftP + ((rightP - leftP) / 2)

	returnLeft := -1

	if nums[leftP] == target {
		returnLeft = leftP
	}

	if leftP+1 == rightP && nums[rightP] == target && nums[leftP] != target {
		return returnLeft + 1
	}

	if leftP+1 == rightP {
		return returnLeft
	}

	if nums[mid] == target && nums[mid-1] < target {
		return mid
	}

	if nums[mid] > target {
		rightP = mid - 1
	} else {
		leftP = mid
	}

	return leftSearch(nums, target, leftP, rightP)
}

func rightSearch(nums []int, target, leftP, rightP int) int {
	mid := leftP + ((rightP - leftP) / 2)
	returnRight := -1

	if nums[rightP] == target {
		returnRight = rightP
	}

	if rightP-1 == leftP && nums[leftP] == target && nums[rightP] != target {
		return returnRight + 1
	}

	if rightP-1 == leftP {
		return returnRight
	}

	if nums[mid] == target && nums[mid+1] > target {
		return mid
	}

	if nums[mid] < target {
		leftP = mid + 1
	} else {
		rightP = mid
	}

	return rightSearch(nums, target, leftP, rightP)
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
```
/*
mainFunction {
  [5, 7, 7, 8, 10]

  // Arguments passed are slice and left and right boundaries
  leftindex = leftSearch([5, 7, 7, 8, 10], left = 0, right = len(slice)-1)
  rightindex = rightSearch([5, 7, 7, 8, 10], left = 0, right = len(slice)-1)

  return []int{leftindex, rightindex}
}

leftSearch() {
  baseCase:
    - mid is equal to target AND value to the left is lesser

  // You either go left or right, with a preference for left
  if mid < target {
    left = mid + 1
  } else {
    right = mid
  }
}

rightSearch() {
  baseCase:
    - mid is equal to target AND value to the right greater

  // You either go left or right, with a preference for left
  if mid > target {
    right = mid - 1
  } else {
    left = mid
  }
}
*/
