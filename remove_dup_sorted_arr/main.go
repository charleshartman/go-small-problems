/* 26. Remove Duplicates from Sorted Array

Given an integer array nums sorted in non-decreasing order, remove the duplicates
in-place such that each unique element appears only once. The relative order of
the elements should be kept the same.

Since it is impossible to change the length of the array in some languages, you
must instead have the result be placed in the first part of the array nums. More
formally, if there are k elements after removing the duplicates, then the first
k elements of nums should hold the final result. It does not matter what you
leave beyond the first k elements.

Return k after placing the final result in the first k slots of nums.

Do not allocate extra space for another array. You must do this by modifying
the input array in-place with O(1) extra memory.

Custom Judge:
The judge will test your solution with the following code:

int[] nums = [...]; // Input array
int[] expectedNums = [...]; // The expected answer with correct length
int k = removeDuplicates(nums); // Calls your implementation
assert k == expectedNums.length;
for (int i = 0; i < k; i++) {
    assert nums[i] == expectedNums[i];
}

If all assertions pass, then your solution will be accepted.

Example 1:
Input: nums = [1,1,2]
Output: 2, nums = [1,2,_]
Explanation: Your function should return k = 2, with the first two elements of nums being 1 and 2 respectively.
It does not matter what you leave beyond the returned k (hence they are underscores).

Example 2:
Input: nums = [0,0,1,1,1,2,2,3,3,4]
Output: 5, nums = [0,1,2,3,4,_,_,_,_,_]
Explanation: Your function should return k = 5, with the first five elements of nums being 0, 1, 2, 3, and 4 respectively.
It does not matter what you leave beyond the returned k (hence they are underscores).

Constraints:
0 <= nums.length <= 3 * 104
-100 <= nums[i] <= 100
nums is sorted in non-decreasing order.
*/

/* PEDAC

UNDERSTAND THE PROBLEM
  Inputs
    - a sorted (non-decreasing order) array of integers that may include duplicates
  Outputs
    - an integer representing the length of the array after duplicates are removed
Rules/Implicit Requirements
  Explicit
    - must mutate input array
    - must not use any additional memory (no hashes)
    - values at the end (after uniques) are arbitrary
    - 0 <= nums.length <= 3 * 104
        - input array may empty, [guard?]
    - -100 <= nums[i] <= 100
  Implicit
    -

EXAMPLES/TEST CASES
  Base Cases
    - Input: nums = [0,0,1,1,1,2,2,3,3,4]
      Output: 5, nums = [0,1,2,3,4,_,_,_,_,_]
    - Input: nums = [1,1,2]
      Output: 2, nums = [1,2,_]
  Edge Cases
    - input array may empty, guard for now

MENTAL MODEL:
  How many pointers do I have?
    2
  Where do they start?
    0 & 0
  When do they move?
    R - moves on every iteration
    W - moves when the value at index R is not equal to the value at index W
  What do I do when each pointer moves?
    - when R moves we start a new iteration
    - when W moves we reassign the value at that index to the value at the index R is at
  When do I cease iteration?
    - when R reaches the boundary of the arr aka len(arr)
      - return the index + 1 of W

ALGORITHM
  - if the slice is of length 0, return 0
  - start r and w at index 0
  - loop over the slice, until r is outside the boundary
    - increment r with every iteration
    - if nums[r] != nums[w]
      - increment w by 1
      - reassign the value to the value at r
*/

package main

import "fmt"

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	r, w := 0, 0

	for r < len(nums) {
		if nums[r] != nums[w] {
			w++
			nums[w] = nums[r]
		}

		r++
	}

	return w + 1
}

func main() {
	s1 := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(removeDuplicates(s1))
	fmt.Println(s1)
	fmt.Println(removeDuplicates([]int{0, 1, 2, 3, 3, 4}))
	fmt.Println(removeDuplicates([]int{1, 1, 2}))
	fmt.Println(removeDuplicates([]int{1, 1}))
	fmt.Println(removeDuplicates([]int{}))
}
