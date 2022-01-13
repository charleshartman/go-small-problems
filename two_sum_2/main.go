/*
167. Two Sum II - Input Array Is Sorted

Given a 1-indexed array of integers numbers that is already sorted in
non-decreasing order, find two numbers such that they add up to a specific
target number. Let these two numbers be numbers[index1] and numbers[index2]
where 1 <= index1 < index2 <= numbers.length.

Return the indices of the two numbers, index1 and index2, added by one as an
integer array [index1, index2] of length 2.

The tests are generated such that there is exactly one solution. You may not
use the same element twice.

Example 1:
Input: numbers = [2,7,11,15], target = 9
Output: [1,2]
Explanation: The sum of 2 and 7 is 9. Therefore, index1 = 1, index2 = 2. We return [1, 2].

Example 2:
Input: numbers = [2,3,4], target = 6
Output: [1,3]
Explanation: The sum of 2 and 4 is 6. Therefore index1 = 1, index2 = 3. We return [1, 3].

Example 3:
Input: numbers = [-1,0], target = -1
Output: [1,2]
Explanation: The sum of -1 and 0 is -1. Therefore index1 = 1, index2 = 2. We return [1, 2].

Constraints:

    2 <= numbers.length <= 3 * 104
    -1000 <= numbers[i] <= 1000
    numbers is sorted in non-decreasing order.
    -1000 <= target <= 1000
    The tests are generated such that there is exactly one solution.



MENTAL MODEL:
  How many pointers do I have?
    2
  Where do they start?
    0 & len(arr) - 1
  When do they move?
    S - moves on when start + end values < target - s++
    E - moves on when start + end values > target - e--
  What do I do when each pointer moves?
    - sum the values and compare with target
  When do I cease iteration?
    - when target number is found return current index values of start and end

Input: numbers = [2,7,11,15], target = 9
Output: [1,2]

algorithm:
func takes a slice of ints and a target int and returns a slice of ints
declare s to 0
declare e to len(sl) - 1
declare currentSum int

perform loop
assign currentSum to sum of s + e
if currentSum equals the target then return new slice {s,e}
elseif currentSum is less than the target then move s (s++)
else currentSum must be greater than target so move e (e--)

*/

package main

import "fmt"

func twoSum(numbers []int, target int) []int {
	s, e := 0, len(numbers)-1
	var currentSum int
	// var result []int

	for s < e {
		currentSum = numbers[s] + numbers[e]
		if currentSum == target {
			return []int{s + 1, e + 1}
			// result = append(result, s+1, e+1)
			// break
		} else if currentSum < target {
			s++
		} else {
			e--
		}
	}

	return nil
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9)) // [1 2]
	fmt.Println(twoSum([]int{2, 3, 4}, 6))      // [1 3]
	fmt.Println(twoSum([]int{-1, 0}, -1))       // [1 2]
	fmt.Println(twoSum([]int{0, 0, 3, 4}, 0))   // [1 2]
	fmt.Println(twoSum([]int{5, 25, 75}, 100))  // [2 3]
}
