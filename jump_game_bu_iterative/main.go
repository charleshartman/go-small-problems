/* 55. Jump Game
You are given an integer array nums. You are initially positioned at the array's
first index, and each element in the array represents your maximum jump length
at that position.

Return true if you can reach the last index, or false otherwise.

Example 1:
Input: nums = [2,3,1,1,4]
Output: true
Explanation: Jump 1 step from index 0 to 1, then 3 steps to the last index.

Example 2:
Input: nums = [3,2,1,0,4]
Output: false
Explanation: You will always arrive at index 3 no matter what. Its maximum jump
length is 0, which makes it impossible to reach the last index.

Constraints:
1 <= nums.length <= 104
0 <= nums[i] <= 105

UNDERSTAND THE PROBLEM
  Inputs
    - an slice on integers, each element has a value of 0 or greater
  Outputs
    - boolean
Rules/Implicit Requirements
  Explicit
    - we are initially positioned at index 0, return true or false based on our
	  ability to "jump" from the starting index to the end of the given slice
    - we may jump from 1 up to the value at the index position we are jumping
	  from (for example given 3, we may jump 1, 2 or 3 positions)
  Implicit
    - when we land on a 0 we may not jump from there
	- a value may enable us to jump "over" the end, this should return true
	- a 0 at the end is possible, we still must jump there to return true

Mental Model
  - iterate through the array backwards
  - check if an element at a position can reach the end within 1..val number of jumps
  - N, number of elements in the slice + sum of the values in the array
  - O(N) or O(N+S)

EXAMPLES/TEST CASES
  Normal Cases
              0 1 2 3 4
    - nums = [2,3,1,1,4] -> true
    - nums = [3,2,1,0,4] -> false

  Edge Cases
    - nums = [3] -> true
    - nums = [2, 5, 0, 0] -> true
    - nums = [0, 1, 0, 1] -> false
    - nums = [1, 0, 1, 0] -> false
    - nums = [3, 0, 8, 2, 0, 0, 1] -> true

DATA STRUCTURE
  Inputs
    - map, key is int, value is bool
  Rules
    - only caching indices with true values

ALGORITHM
  	- create a map (dp)
	- add the last value from nums into the cache (dp) with a value of true
  	- outer loop, iterates over every element in nums except the last (moving backwards)
    	- Identify the value at the current index in nums
    	- inner loop, iterate from 1 to the value at the current element from nums
      		- if dp has a value for the current index + i
        	- add the key for the current index with a true value
        	- and break
  - return dp[0]

CODE
*/

package main

import "fmt"

func canJump(nums []int) bool {
	// Initialize cache to hold index positions that can successfully jump to
	// the final element (queries to keys that do not exist in the map will
	// implicitly return false)
	dp := map[int]bool{}

	// Last index is true because we already reached the end of the slice
	dp[len(nums)-1] = true

	// Iterate through every element in the array starting from the last up
	// to the first
	for i := len(nums) - 2; i >= 0; i-- {
		val := nums[i]
		// Iterate from 1 up to the value of the current element
		for j := 1; j <= val; j++ {
			// If dp[i + j] == true, this means that we can jump from our
			// current index to an index that can reach the end of the slice.
			if dp[i+j] == true {
				dp[i] = true
				break
			}
		}
	}
	// After all iterations we return the value in our map for the first index,
	// this will be true if the jump is possible, or absent (thus false) if not
	return dp[0]
}

func main() {
	fmt.Println(canJump([]int{2, 3, 1, 1, 4}))       // true
	fmt.Println(canJump([]int{3, 2, 1, 0, 4}))       // false
	fmt.Println(canJump([]int{3, 0, 8, 2, 0, 0, 1})) // true
	fmt.Println(canJump([]int{1, 0, 1, 0}))          // false
	fmt.Println(canJump([]int{2, 5, 0, 0}))          // true
	fmt.Println(canJump([]int{0, 1, 0, 1}))          // false
}
