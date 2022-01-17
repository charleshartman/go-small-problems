/* 234. Palindrome Linked List

Given the head of a singly linked list, return true if it is a palindrome.

Example 1:
Input: head = [1,2,2,1]
Output: true

Example 2:
Input: head = [1,2]
Output: false

Constraints:
The number of nodes in the list is in the range [1, 105].
0 <= Node.val <= 9

Follow up: Could you do it in O(n) time and O(1) space?
*/

/* algorithm 1
- brute force-ish... O(N) space complexity
- O(N) time complexity O(2N) drop constant O(N)
	- traverse linked list and insert each value into a slice
	- use helper function to determine (with pointers) if palindrome
		- start and end var, 0 and len(sl) - 1
		- iterate checking equality, if != return false, else true
	- return result of calling helper function
*/

package main

import (
	"fmt"

	"github.com/charleshartman/tortoise"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// func isSlicePalindrome(sl []int) bool {
// 	s, e := 0, len(sl)-1
// 	result := true
// 	for s <= e {
// 		if sl[s] != sl[e] {
// 			result = false
// 			break
// 		}
// 		s++
// 		e--
// 	}

// 	return result
// }

func isPalindrome(head *ListNode) bool {
	values := []int{}
	dummy := new(ListNode)
	dummy.Next = head
	r := dummy.Next
	for r != nil {
		values = append(values, r.Val)
		r = r.Next
	}

	return tortoise.IsIntSlicePalindrome(values)
}

func main() {
	// fmt.Println(isSlicePalindrome([]int{1, 2, 2, 1}))    // true
	// fmt.Println(isSlicePalindrome([]int{1, 3, 2, 1}))    // false
	// fmt.Println(isSlicePalindrome([]int{1, 2, 3, 2, 1})) // true
	testHead := new(ListNode)
	testHead.Val = 1
	testElement1 := new(ListNode)
	testElement1.Val = 2
	testHead.Next = testElement1
	testElement2 := new(ListNode)
	testElement2.Val = 2
	testElement1.Next = testElement2
	testElement3 := new(ListNode)
	testElement3.Val = 1
	testElement2.Next = testElement3
	//
	fmt.Println(isPalindrome(testHead))
}
