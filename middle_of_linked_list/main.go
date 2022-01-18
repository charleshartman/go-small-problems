/* 876. Middle of the Linked List

Given the head of a singly linked list, return the middle node of the linked list.

If there are two middle nodes, return the second middle node.

Example 1:
Input: head = [1,2,3,4,5]
Output: [3,4,5]
Explanation: The middle node of the list is node 3.

Example 2:
Input: head = [1,2,3,4,5,6]
Output: [4,5,6]
Explanation: Since the list has two middle nodes with values 3 and 4, we return the second one.

Constraints:
The number of nodes in the list is in the range [1, 100].
1 <= Node.val <= 100
*/

/* algorithm
traverse the list with runner, building length counter
traverse the list with anchor to the nth node where n is the counter / 2 rounded up
return the node at anchor

*/

package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// func middleNode(head *ListNode) *ListNode {
// 	a, r := head, head
// 	nth := 1
// 	for r.Next != nil {
// 		r = r.Next
// 		nth++
// 	}

// 	for i := 1; i <= nth/2; i++ {
// 		a = a.Next
// 	}

// 	fmt.Println(nth, a, r)
// 	return a
// }

func middleNode(head *ListNode) *ListNode {
	a, r := head, head
	for r != nil && r.Next != nil {
		r = r.Next.Next
		a = a.Next
	}

	return a
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
	fmt.Println(middleNode(testHead))
}
