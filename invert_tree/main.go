/*
Given the root of a binary tree, invert the tree, and return its root.

Example 1:
Input: root = [4,2,7,1,3,6,9]
Output: [4,7,2,9,6,3,1]

Example 2:
Input: root = [2,1,3]
Output: [2,3,1]

Example 3:
Input: root = []
Output: []

Constraints:
The number of nodes in the tree is in the range [0, 100].
-100 <= Node.val <= 100

*/

/*
func someProblem(root *treeNode) *treeNode {
  // base case
  left, right := root.Left, root.Right
  leftResult := someProblem(left)
  rightResult := someProblem(right)

  result := // combine root, leftResult and rightResult
  return result
}
*/

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)
	return root

	// root's left node should be the inverted tree of the root's right node
	// root's right node should be the inverted tree of the root's left node
}
