/* 206. Reverse Linked List

Given the head of a singly linked list, reverse the list, and return
the reversed list.

Example 1:
Input: head = [1,2,3,4,5]
Output: [5,4,3,2,1]

Example 2:
Input: head = [1,2]
Output: [2,1]

Example 3:
Input: head = []
Output: []

Constraints:
The number of nodes in the list is the range [0, 5000].
-5000 <= Node.val <= 5000

*/

/* algorithm

 */

package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	var p *ListNode
	c := head
	n := head.Next

	for n != nil {
		c.Next = p
		p = c
		c = n
		n = n.Next
	}

	c.Next = p
	head = c
	return head
}

func main() {
	testHead := new(ListNode)
	testHead.Val = 1
	testElement1 := new(ListNode)
	testElement1.Val = 2
	testHead.Next = testElement1
	testElement2 := new(ListNode)
	testElement2.Val = 3
	testElement1.Next = testElement2
	testElement3 := new(ListNode)
	testElement3.Val = 4
	testElement2.Next = testElement3
	//
	fmt.Println(reverseList(testHead))
}
