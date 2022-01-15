/* 21. Merge Two Sorted Lists

You are given the heads of two sorted linked lists list1 and list2.

Merge the two lists in a one sorted list. The list should be made by splicing
together the nodes of the first two lists.

Return the head of the merged linked list.

Example 1:
Input: list1 = [1,2,4], list2 = [1,3,4]
Output: [1,1,2,3,4,4]

Example 2:
Input: list1 = [], list2 = []
Output: []

Example 3:
Input: list1 = [], list2 = [0]
Output: [0]

Constraints:

    The number of nodes in both lists is in the range [0, 50].
    -100 <= Node.val <= 100
    Both list1 and list2 are sorted in non-decreasing order.

template for 3 pointer slide for LL

func something(head *ListNode) *ListNode {
  if head == nil || head.Next == nil {
    return head
  }
  dummy := new(ListNode)
  dummy.Next = head
  prev, curr, next := dummy, head, head.Next

  for next != nil {
    // do some comparison
    // adjust pointers
  }
  // return new head for in place solution
}

*/

/* ALGORITHM
- declare dummy
- declare head to dummy.Next
- while list1 isn't nil and list2 isn't nil
  - if list1 node value is greater than list2 node value

- return dummy.head.Next
*/
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummyHead := new(ListNode)
	current := dummyHead

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			current.Next = list1
			list1 = list1.Next
		} else {
			current.Next = list2
			list2 = list2.Next
		}
		current = current.Next
	}

	if list1 != nil {
		current.Next = list1
	}

	if list2 != nil {
		current.Next = list2
	}

	return dummyHead.Next
}
