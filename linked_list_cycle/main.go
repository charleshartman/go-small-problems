/*
UNDERSTAND THE PROBLEM
  Inputs
    - head of a linked list
  Outputs
    - boolean, true if there is a cycle
Rules/Implicit Requirements
  Explicit
    - find a cycle, if we can reach the same node again (previously seen node) then true
    - traverse the link list forward
  Implicit
    -
Clarifying Questions
  - what happens with one node? can it loop to itself?
Mental Model (optional)
  - We know there is a cycle (thus return true) if the fast pointer meets with the slow pointer
  - We know to return false if fastptr reaches nil

EXAMPLES/TEST CASES
  Base Cases
  Example 1:
  Input: head = [3,2,0,-4], pos = 1
  Output: true
  Explanation: There is a cycle in the linked list, where the tail connects to the 1st node (0-indexed).

  Example 2:
  Input: head = [1,2], pos = 0
  Output: true
  Explanation: There is a cycle in the linked list, where the tail connects to the 0th node.

  Example 3:
  Input: head = [1], pos = -1
  Output: false
  Explanation: There is no cycle in the linked list.

  Edge Cases
    - Single node linked list and (pos) equals 0
    - Empty node (pos) == -1

DATA STRUCTURE
  - Head node from linked list

ALGORITHM
  - Given a head node
  - add guard clause to account for nil head node
  - Define two pointers that both point to the head node (s and f)

  - Loop through linked list until the fast pointer (f) reaches nil
    - f moves two nodes at a time:
      1. f moves one node and we check if the next value is nil
        - If it is nil, break
      2. Check if node that points at s is the same as f
        - If so, return true
      3. f moves one node and we check if the next value is nil
        - If it is nil, break
      4. Check if node that points at s is the same as f
        - If so, return true
    - s moves one node at a time

  - Return false

  head = [3,2,0,-4], pos = 1
            S
            F
  Input: head = [1], pos = -1

CODE
*/

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	s, f := head, head

	for f != nil {
		for i := 0; i < 2; i++ {
			f = f.Next
			if f == nil {
				return false
			}
			if s == f {
				return true
			}
		}

		s = s.Next
	}

	return false
}
