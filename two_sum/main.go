/*
Come up with, and propose the naive algorithm. This is a way to wrap your head around the problem and verify with the interviewer that you've got the correct understanding of the problem
Go down the Big O curves one at a time and see if you can improve the efficiency of your solution
Use of data structures?
Extra space / in place?
For array problems, this is very often a process to go from O(N^2) to O(N)

Option 1: Naive: nested loops and compare - T: O(N^2), S: O(1)
- nested loop, outer loop (i) and inner loop (j) both iterate over the same array
- check that i != j && the arr[i] + arr[j] = target
  - if that's true then return the indices i, j in an array
- otherwise, you've iterated through the nested loop and didn't find a pair
  - return []

Option 2: Sort and Binary Search for the complementary number - O(N Log N), S: O(1)
- sort the array => O(NLogN)
- for each number in the array, perform a binary search for the counterpart
- O(2*NLogN) => O(NLogN)

Option 3: Hash table - T: O(N), S: O(N)
- counterpart = target - arr[i]
- looking up the counterpart in a ref. hash
  - if it's not there, then add arr[i] to the hash w/ index
  - if it's there then just return the current index i with the index of the counterpart (value corresponding to the key)

Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

You can return the answer in any order.

Example 1:

Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
Example 2:

Input: nums = [3,2,4], target = 6
Output: [1,2]
Example 3:

Input: nums = [3,3], target = 6
Output: [0,1]


Constraints:

2 <= nums.length <= 104
-109 <= nums[i] <= 109
-109 <= target <= 109
Only one valid answer exists.


Follow-up: Can you come up with an algorithm that is less than O(n2) time complexity?
*/

package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9)) // [0, 1]
	fmt.Println(twoSum([]int{3, 2, 4}, 6))      // [1, 2]
	fmt.Println(twoSum([]int{3, 3}, 6))         // [0, 1]
}

func twoSum(s []int, t int) []int {
	seen := map[int]int{}
	var result []int

	for i, num := range s {
		c := t - num

		if val, ok := seen[c]; ok {
			result = []int{i, val}
			break
		}

		seen[num] = i
	}

	return result
}
