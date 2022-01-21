/* 114. Flatten Binary Tree to Linked List

Given the root of a binary tree, flatten the tree into a "linked list":

The "linked list" should use the same TreeNode class where the right child
pointer points to the next node in the list and the left child pointer is always null.

The "linked list" should be in the same order as a pre-order traversal of the binary tree.

Example 1:
Input: root = [1,2,5,3,4,null,6]
Output: [1,null,2,null,3,null,4,null,5,null,6]

Example 2:
Input: root = []
Output: []

Example 3:
Input: root = [0]
Output: [0]

Constraints:
The number of nodes in the tree is in the range [0, 2000].
-100 <= Node.val <= 100

Follow up: Can you flatten the tree in-place (with O(1) extra space)?
*/

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) (result []*TreeNode) {
	if root == nil {
		return
	}

	result = append(result, root)
	result = append(result, preorderTraversal(root.Left)...)
	result = append(result, preorderTraversal(root.Right)...)

	return
}

func flatten(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	orderedNodes := preorderTraversal(root) // slice of TreeNodes in pre-order

	for i := 0; i < len(orderedNodes)-1; i++ {
		orderedNodes[i].Right = orderedNodes[i+1]
		orderedNodes[i].Left = nil
	}

	return orderedNodes[0]
}
