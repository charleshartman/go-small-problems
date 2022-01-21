/*
Given the root of a binary tree, return the level order traversal of its
nodes' values. (i.e., from left to right, level by level).

Example 1:
Input: root = [3,9,20,null,null,15,7]
Output: [[3],[9,20],[15,7]]

Example 2:
Input: root = [1]
Output: [[1]]

Example 3:
Input: root = []
Output: []

Algorithm
printLevelorder(tree)
1) Create an empty queue q
2) temp_node = root /*start from root
3) Loop while temp_node is not NULL
    a) print temp_node->data. (push into result array/slice)
    b) Enqueue temp_node’s children
      (first left then right children) to q
    c) Dequeue a node from q. (temp_node = dequeued node)
*/

package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) (result []int) {
	q := Queue{}
	tempNode := root

	for tempNode != nil {
		result = append(result, tempNode.Val)
		if tempNode.Left != nil {
			q.Enqueue(tempNode.Left)
		}
		if tempNode.Right != nil {
			q.Enqueue(tempNode.Right)
		}

		tempNode = q.Dequeue()
	}
	return
}

type Queue struct {
	list []*TreeNode
}

func (q *Queue) Enqueue(value *TreeNode) {
	q.list = append(q.list, value)
}

func (q *Queue) Dequeue() *TreeNode {
	if len(q.list) == 0 {
		return nil
	}
	removed := q.list[0]
	q.list = q.list[1:]
	return removed
}

func main() {
	// root = [3,9,20,null,null,15,7]
	node5 := &TreeNode{7, nil, nil}
	node4 := &TreeNode{15, nil, nil}
	node3 := &TreeNode{20, node4, node5}
	node2 := &TreeNode{9, nil, nil}
	node1 := &TreeNode{3, node2, node3}
	fmt.Println(levelOrder(node1))
}
