package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	f := head
	l := head

	for i := 0; i < n; i++ {
		l = l.Next
	}

	if l == nil {
		head = f.Next
		return head
	}

	for l.Next != nil {
		f = f.Next
		l = l.Next
	}

	f.Next = f.Next.Next

	return head
}
